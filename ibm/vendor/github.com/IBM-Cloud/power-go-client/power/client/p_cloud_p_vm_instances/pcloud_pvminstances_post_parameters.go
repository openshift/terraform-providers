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
	"github.com/go-openapi/swag"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudPvminstancesPostParams creates a new PcloudPvminstancesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesPostParams() *PcloudPvminstancesPostParams {
	return &PcloudPvminstancesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesPostParamsWithTimeout creates a new PcloudPvminstancesPostParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesPostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesPostParams {
	return &PcloudPvminstancesPostParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesPostParamsWithContext creates a new PcloudPvminstancesPostParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesPostParamsWithContext(ctx context.Context) *PcloudPvminstancesPostParams {
	return &PcloudPvminstancesPostParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesPostParamsWithHTTPClient creates a new PcloudPvminstancesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesPostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesPostParams {
	return &PcloudPvminstancesPostParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesPostParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances post operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesPostParams struct {

	/* Body.

	   Parameters for the creation of a new Power VM Instance
	*/
	Body *models.PVMInstanceCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* SkipHostValidation.

	   Option to skip host validation on PVMInstance Create API
	*/
	SkipHostValidation *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesPostParams) WithDefaults() *PcloudPvminstancesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithContext(ctx context.Context) *PcloudPvminstancesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithBody(body *models.PVMInstanceCreate) *PcloudPvminstancesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetBody(body *models.PVMInstanceCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithSkipHostValidation adds the skipHostValidation to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) WithSkipHostValidation(skipHostValidation *bool) *PcloudPvminstancesPostParams {
	o.SetSkipHostValidation(skipHostValidation)
	return o
}

// SetSkipHostValidation adds the skipHostValidation to the pcloud pvminstances post params
func (o *PcloudPvminstancesPostParams) SetSkipHostValidation(skipHostValidation *bool) {
	o.SkipHostValidation = skipHostValidation
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SkipHostValidation != nil {

		// query param skipHostValidation
		var qrSkipHostValidation bool

		if o.SkipHostValidation != nil {
			qrSkipHostValidation = *o.SkipHostValidation
		}
		qSkipHostValidation := swag.FormatBool(qrSkipHostValidation)
		if qSkipHostValidation != "" {

			if err := r.SetQueryParam("skipHostValidation", qSkipHostValidation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
