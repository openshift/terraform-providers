// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesSnapshotsRestorePostParams creates a new PcloudPvminstancesSnapshotsRestorePostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesSnapshotsRestorePostParams() *PcloudPvminstancesSnapshotsRestorePostParams {
	return &PcloudPvminstancesSnapshotsRestorePostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesSnapshotsRestorePostParamsWithTimeout creates a new PcloudPvminstancesSnapshotsRestorePostParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesSnapshotsRestorePostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesSnapshotsRestorePostParams {
	return &PcloudPvminstancesSnapshotsRestorePostParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesSnapshotsRestorePostParamsWithContext creates a new PcloudPvminstancesSnapshotsRestorePostParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesSnapshotsRestorePostParamsWithContext(ctx context.Context) *PcloudPvminstancesSnapshotsRestorePostParams {
	return &PcloudPvminstancesSnapshotsRestorePostParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesSnapshotsRestorePostParamsWithHTTPClient creates a new PcloudPvminstancesSnapshotsRestorePostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesSnapshotsRestorePostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesSnapshotsRestorePostParams {
	return &PcloudPvminstancesSnapshotsRestorePostParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesSnapshotsRestorePostParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances snapshots restore post operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesSnapshotsRestorePostParams struct {

	/* Body.

	   PVM Instance snapshot restore parameters
	*/
	Body *models.SnapshotRestore

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	/* RestoreFailAction.

	   Action to take on a failed snapshot restore
	*/
	RestoreFailAction *string

	/* SnapshotID.

	   PVM Instance snapshot id
	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances snapshots restore post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithDefaults() *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances snapshots restore post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithContext(ctx context.Context) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithBody(body *models.SnapshotRestore) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetBody(body *models.SnapshotRestore) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WithRestoreFailAction adds the restoreFailAction to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithRestoreFailAction(restoreFailAction *string) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetRestoreFailAction(restoreFailAction)
	return o
}

// SetRestoreFailAction adds the restoreFailAction to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetRestoreFailAction(restoreFailAction *string) {
	o.RestoreFailAction = restoreFailAction
}

// WithSnapshotID adds the snapshotID to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WithSnapshotID(snapshotID string) *PcloudPvminstancesSnapshotsRestorePostParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the pcloud pvminstances snapshots restore post params
func (o *PcloudPvminstancesSnapshotsRestorePostParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesSnapshotsRestorePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if o.RestoreFailAction != nil {

		// query param restore_fail_action
		var qrRestoreFailAction string

		if o.RestoreFailAction != nil {
			qrRestoreFailAction = *o.RestoreFailAction
		}
		qRestoreFailAction := qrRestoreFailAction
		if qRestoreFailAction != "" {

			if err := r.SetQueryParam("restore_fail_action", qRestoreFailAction); err != nil {
				return err
			}
		}
	}

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
