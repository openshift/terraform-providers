package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pkg/errors"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/govc/importx"
	"github.com/vmware/govmomi/nfc"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/ovf"
	"github.com/vmware/govmomi/vapi/tags"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/soap"
	"github.com/vmware/govmomi/vim25/types"
)

func resourceVSpherePrivateImportOva() *schema.Resource {
	return &schema.Resource{
		Create:        resourceVSpherePrivateImportOvaCreate,
		Read:          resourceVSpherePrivateImportOvaRead,
		Delete:        resourceVSpherePrivateImportOvaDelete,
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of the virtual machine that will be created.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"filename": {
				Type:         schema.TypeString,
				Description:  "The filename path to the ova file to be imported.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"datacenter": {
				Type:         schema.TypeString,
				Description:  "The name of the datacenter.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"cluster": {
				Type:         schema.TypeString,
				Description:  "The name of the cluster.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"resource_pool": {
				Type:         schema.TypeString,
				Description:  "The absolute path to the resource pool.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"network": {
				Type:         schema.TypeString,
				Description:  "The name of a network that the virtual machine will use.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"datastore": {
				Type:         schema.TypeString,
				Description:  "The name of the virtual machine's datastore.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"folder": {
				Type:         schema.TypeString,
				Description:  "The name of the folder to locate the virtual machine in.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"tag": {
				Type:         schema.TypeString,
				Description:  "The name of the tag to attach the virtual machine in.",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"disk_type": {
				Type:        schema.TypeString,
				Description: "The name of the disk provisioning type, valid values are thin, thick and eagerZeroedThick.",
				Optional:    true,
				ForceNew:    true,
			},
		},
	}
}

// importOvaParams contains the vCenter objects required to import a OVA into vSphere.
type importOvaParams struct {
	ResourcePool    *object.ResourcePool
	Datacenter      *object.Datacenter
	Datastore       *object.Datastore
	Network         *object.Network
	Host            *object.HostSystem
	Folder          *object.Folder
	HardwareVersion string
}

func findImportOvaParams(client *vim25.Client, datacenter, cluster, resourcePool, datastore, network, folder string) (*importOvaParams, error) {

	var ccrMo mo.ClusterComputeResource
	var folderPath string

	ctx, cancel := context.WithTimeout(context.TODO(), defaultAPITimeout)
	defer cancel()

	importOvaParams := &importOvaParams{}
	finder := find.NewFinder(client)

	// Find the object Datacenter by using its name provided by install-config
	dcObj, err := finder.Datacenter(ctx, datacenter)
	if err != nil {
		return nil, err
	}
	importOvaParams.Datacenter = dcObj

	// When finder.Datacenter is executed apparently it
	// does not set the datacenter, why, who knows
	// Replace finder with finder.SetDatacenter
	finder = finder.SetDatacenter(dcObj)

	folderPath = folder
	if !strings.HasPrefix(folder, "/") {
		folderPath = fmt.Sprintf("/%s/vm/%s", datacenter, folder)
	}

	// Create an absolute path to the folder in case the provided folder is nested.
	folderObj, err := finder.Folder(ctx, folderPath)
	if err != nil {
		return nil, err
	}
	importOvaParams.Folder = folderObj

	// Find the resource pool object by using its path provided by install-config,
	// or generated in pkg/asset/machines/vsphere/machines.go
	resourcePoolObj, err := finder.ResourcePool(ctx, resourcePool)
	if err != nil {
		return nil, err
	}
	importOvaParams.ResourcePool = resourcePoolObj

	clusterPath := cluster
	if !strings.HasPrefix(cluster, "/") {
		clusterPath = fmt.Sprintf("/%s/host/%s", datacenter, cluster)
	}

	clusterComputeResource, err := finder.ClusterComputeResource(ctx, clusterPath)
	if err != nil {
		return nil, err
	}

	// Get the network properties that is defined in ClusterComputeResource
	// We need to know if the network name provided exists in the cluster that was
	// also provided.
	err = clusterComputeResource.Properties(ctx, clusterComputeResource.Reference(), []string{"network"}, &ccrMo)
	if err != nil {
		return nil, err
	}

	// Find the network object in the list of networks inside cluster
	for _, networkMoRef := range ccrMo.Network {
		if networkMoRef.Value == network {
			importOvaParams.Network = object.NewNetwork(client, networkMoRef)
			break
		}
	}

	if importOvaParams.Network == nil {
		return nil, errors.Errorf("failed to find a host in the cluster that contains the provided network")
	}

	// Find all the datastores that are configured under the cluster
	clusterDatastores, err := clusterComputeResource.Datastores(ctx)
	if err != nil {
		return nil, err
	}

	if len(clusterDatastores) == 0 {
		return nil, errors.Errorf("failed to find any datastore(s) in the cluster")
	}

	// Find the specific datastore by the name provided
	for _, objectDS := range clusterDatastores {

		objectDSName, err := objectDS.ObjectName(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "unable to find datastore name")
		}
		// clusterComputeResource.Datastores(ctx) does not properly
		// handle Common - which includes InventoryPath and .Name()
		// To workaround this issue we must retrieve the Datastore
		// object again with the method below. Do not remove this.
		ds, err := finder.Datastore(ctx, objectDSName)

		if err != nil {
			return nil, errors.Wrap(err, "unable to find datastore object by name")
		}
		datastoreName := ds.Name()
		datastorePath := ds.InventoryPath

		if datastoreName == datastore || datastorePath == datastore {
			importOvaParams.Datastore = ds
			break
		}
	}
	if importOvaParams.Datastore == nil {
		return nil, errors.Errorf("failed to find a host in the cluster that contains the provided datastore")
	}

	// Find all the HostSystem(s) under cluster
	hosts, err := clusterComputeResource.Hosts(ctx)
	if err != nil {
		return nil, err
	}
	var hostSystemManagedObject mo.HostSystem

	// Confirm that the network and datastore that was provided is
	// available for use on the HostSystem we will import the
	// OVA to.
	for _, hostObj := range hosts {

		foundDatastore := false
		foundNetwork := false
		err := hostObj.Properties(ctx, hostObj.Reference(), []string{"config.product", "network", "datastore", "runtime"}, &hostSystemManagedObject)
		if err != nil {
			return nil, err
		}

		// Skip all hosts that are in maintenance mode.
		if hostSystemManagedObject.Runtime.InMaintenanceMode {
			continue
		}
		for _, dsMoRef := range hostSystemManagedObject.Datastore {

			if importOvaParams.Datastore.Reference().Value == dsMoRef.Value {
				foundDatastore = true
				break
			}
		}
		for _, nMoRef := range hostSystemManagedObject.Network {
			if importOvaParams.Network.Reference().Value == nMoRef.Value {
				foundNetwork = true
				break
			}
		}

		if foundDatastore && foundNetwork {
			return importOvaParams, nil
		}
	}

	if importOvaParams.Datastore != nil && importOvaParams.Network != nil {
		return importOvaParams, nil
	}

	return nil, errors.Errorf("failed to find a host in the cluster that contains the provided datastore and network")
}

func attachTag(d *schema.ResourceData, meta interface{}) error {
	ctx, cancel := context.WithTimeout(context.TODO(), defaultAPITimeout)
	defer cancel()
	tagManager := tags.NewManager(meta.(*VSphereClient).restClient)
	moRef := types.ManagedObjectReference{
		Value: d.Id(),
		Type:  "VirtualMachine",
	}

	err := tagManager.AttachTag(ctx, d.Get("tag").(string), moRef)

	if err != nil {
		return err
	}
	return nil
}

// Used govc/importx/ovf.go as an example to implement
// resourceVspherePrivateImportOvaCreate and upload functions
// See: https://github.com/vmware/govmomi/blob/cc10a0758d5b4d4873388bcea417251d1ad03e42/govc/importx/ovf.go#L196-L324
func upload(ctx context.Context, archive *importx.ArchiveFlag, lease *nfc.Lease, item nfc.FileItem) error {
	file := item.Path

	f, size, err := archive.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	opts := soap.Upload{
		ContentLength: size,
	}

	return lease.Upload(ctx, item, f, opts)
}

func resourceVSpherePrivateImportOvaCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: Beginning import ova create", d.Get("filename").(string))

	ctx, cancel := context.WithTimeout(context.TODO(), defaultAPITimeout)
	defer cancel()
	client := meta.(*VSphereClient).vimClient.Client
	filename := d.Get("filename").(string)
	archive := &importx.ArchiveFlag{Archive: &importx.TapeArchive{Path: filename}}

	importOvaParams, err := findImportOvaParams(client,
		d.Get("datacenter").(string),
		d.Get("cluster").(string),
		d.Get("resource_pool").(string),
		d.Get("datastore").(string),
		d.Get("network").(string),
		d.Get("folder").(string))
	if err != nil {
		return errors.Errorf("failed to find provided vSphere objects: %s", err)
	}

	ovfDescriptor, err := archive.ReadOvf("*.ovf")
	if err != nil {
		// Open the corrupt OVA file
		f, ferr := os.Open(filename)
		if ferr != nil {
			err = fmt.Errorf("%s, %w", err.Error(), ferr)
		}
		defer f.Close()

		// Get a sha256 on the corrupt OVA file
		// and the size of the file
		h := sha256.New()
		written, cerr := io.Copy(h, f)
		if cerr != nil {
			err = fmt.Errorf("%s, %w", err.Error(), cerr)
		}

		return errors.Errorf("ova %s has a sha256 of %x and a size of %d bytes, failed to read the ovf descriptor %s", filename, h.Sum(nil), written, err)
	}

	ovfEnvelope, err := archive.ReadEnvelope(ovfDescriptor)
	if err != nil {
		return errors.Errorf("failed to parse ovf: %s", err)
	}

	// The RHCOS OVA only has one network defined by default
	// The OVF envelope defines this.  We need a 1:1 mapping
	// between networks with the OVF and the host
	if len(ovfEnvelope.Network.Networks) != 1 {
		return errors.Errorf("Expected the OVA to only have a single network adapter")
	}
	// Create mapping between OVF and the network object
	// found by Name
	networkMappings := []types.OvfNetworkMapping{{
		Name:    ovfEnvelope.Network.Networks[0].Name,
		Network: importOvaParams.Network.Reference(),
	}}

	// This is a very minimal spec for importing an OVF.
	cisp := types.OvfCreateImportSpecParams{
		EntityName:     d.Get("name").(string),
		NetworkMapping: networkMappings,
	}

	switch diskType := d.Get("disk_type"); diskType {
	case "":
		// Disk provisioning type will be set according to the default storage policy of vsphere.
	case "thin":
		cisp.DiskProvisioning = string(types.OvfCreateImportSpecParamsDiskProvisioningTypeThin)
	case "thick":
		cisp.DiskProvisioning = string(types.OvfCreateImportSpecParamsDiskProvisioningTypeThick)
	case "eagerZeroedThick":
		cisp.DiskProvisioning = string(types.OvfCreateImportSpecParamsDiskProvisioningTypeEagerZeroedThick)
	default:
		return errors.Errorf("Disk provisioning type %q is not supported.", diskType)
	}

	m := ovf.NewManager(client)
	spec, err := m.CreateImportSpec(ctx,
		string(ovfDescriptor),
		importOvaParams.ResourcePool.Reference(),
		importOvaParams.Datastore.Reference(),
		cisp)

	if err != nil {
		return errors.Errorf("failed to create import spec: %s", err)
	}
	if spec.Error != nil {
		return errors.New(spec.Error[0].LocalizedMessage)
	}

	// The lease and upload cannot be used with a timeout
	// since we do not know how long it will take to upload
	// the ova to vSphere
	ctx = context.TODO()

	//Creates a new entity in this resource pool.
	//See VMware vCenter API documentation: Managed Object - ResourcePool - ImportVApp
	lease, err := importOvaParams.ResourcePool.ImportVApp(ctx,
		spec.ImportSpec,
		importOvaParams.Folder,
		importOvaParams.Host)

	if err != nil {
		return errors.Errorf("failed to import vapp: %s", err)
	}

	info, err := lease.Wait(ctx, spec.FileItem)
	if err != nil {
		return errors.Errorf("failed to lease wait: %s", err)
	}

	d.SetId(info.Entity.Value)

	err = attachTag(d, meta)
	if err != nil {
		return errors.Errorf("failed to attach tag to virtual machine: %s", err)
	}

	u := lease.StartUpdater(ctx, info)
	defer u.Done()

	for _, i := range info.Items {
		// upload the vmdk to which ever host that was first
		// available with the required network and datastore.
		err = upload(ctx, archive, lease, i)
		if err != nil {
			return errors.Errorf("failed to upload: %s", err)
		}
	}

	err = lease.Complete(ctx)
	if err != nil {
		return errors.Errorf("failed to lease complete: %s", err)
	}

	log.Printf("[DEBUG] %s: ova import complete", d.Get("name").(string))

	vm := object.NewVirtualMachine(client, info.Entity)
	if vm == nil {
		return fmt.Errorf("error VirtualMachine not found, managed object id: %s", d.Id())
	}
	log.Printf("[DEBUG] %s: mark as template", vm.Name())

	err = vm.MarkAsTemplate(ctx)
	if err != nil {
		return errors.Errorf("failed to mark vm as template: %s", err)
	}
	log.Printf("[DEBUG] %s: mark as template complete", vm.Name())

	return resourceVSpherePrivateImportOvaRead(d, meta)
}

func resourceVSpherePrivateImportOvaRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*VSphereClient).vimClient.Client
	moRef := types.ManagedObjectReference{
		Value: d.Id(),
		Type:  "VirtualMachine",
	}

	vm := object.NewVirtualMachine(client, moRef)
	if vm == nil {
		return fmt.Errorf("error VirtualMachine not found, managed object id: %s", d.Id())
	}

	return nil
}

func resourceVSpherePrivateImportOvaDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] %s: Beginning delete", d.Get("name").(string))
	ctx, cancel := context.WithTimeout(context.TODO(), defaultAPITimeout)
	defer cancel()

	client := meta.(*VSphereClient).vimClient.Client
	moRef := types.ManagedObjectReference{
		Value: d.Id(),
		Type:  "VirtualMachine",
	}

	vm := object.NewVirtualMachine(client, moRef)
	if vm == nil {
		return errors.Errorf("VirtualMachine not found")
	}

	task, err := vm.Destroy(ctx)
	if err != nil {
		return errors.Errorf("failed to destroy virtual machine %s", err)
	}

	err = task.Wait(ctx)
	if err != nil {
		return errors.Errorf("failed to destroy virtual machine %s", err)
	}

	d.SetId("")

	log.Printf("[DEBUG] %s: Delete complete", d.Get("name").(string))

	return nil
}
