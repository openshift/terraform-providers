// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIkepoliciesDeleteReader is a Reader for the PcloudIkepoliciesDelete structure.
type PcloudIkepoliciesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIkepoliciesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudIkepoliciesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudIkepoliciesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudIkepoliciesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudIkepoliciesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudIkepoliciesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudIkepoliciesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudIkepoliciesDeleteOK creates a PcloudIkepoliciesDeleteOK with default headers values
func NewPcloudIkepoliciesDeleteOK() *PcloudIkepoliciesDeleteOK {
	return &PcloudIkepoliciesDeleteOK{}
}

/*
PcloudIkepoliciesDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudIkepoliciesDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud ikepolicies delete o k response has a 2xx status code
func (o *PcloudIkepoliciesDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud ikepolicies delete o k response has a 3xx status code
func (o *PcloudIkepoliciesDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete o k response has a 4xx status code
func (o *PcloudIkepoliciesDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud ikepolicies delete o k response has a 5xx status code
func (o *PcloudIkepoliciesDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud ikepolicies delete o k response a status code equal to that given
func (o *PcloudIkepoliciesDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudIkepoliciesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudIkepoliciesDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteOK  %+v", 200, o.Payload)
}

func (o *PcloudIkepoliciesDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesDeleteBadRequest creates a PcloudIkepoliciesDeleteBadRequest with default headers values
func NewPcloudIkepoliciesDeleteBadRequest() *PcloudIkepoliciesDeleteBadRequest {
	return &PcloudIkepoliciesDeleteBadRequest{}
}

/*
PcloudIkepoliciesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudIkepoliciesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud ikepolicies delete bad request response has a 2xx status code
func (o *PcloudIkepoliciesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud ikepolicies delete bad request response has a 3xx status code
func (o *PcloudIkepoliciesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete bad request response has a 4xx status code
func (o *PcloudIkepoliciesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud ikepolicies delete bad request response has a 5xx status code
func (o *PcloudIkepoliciesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud ikepolicies delete bad request response a status code equal to that given
func (o *PcloudIkepoliciesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudIkepoliciesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudIkepoliciesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudIkepoliciesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesDeleteUnauthorized creates a PcloudIkepoliciesDeleteUnauthorized with default headers values
func NewPcloudIkepoliciesDeleteUnauthorized() *PcloudIkepoliciesDeleteUnauthorized {
	return &PcloudIkepoliciesDeleteUnauthorized{}
}

/*
PcloudIkepoliciesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudIkepoliciesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud ikepolicies delete unauthorized response has a 2xx status code
func (o *PcloudIkepoliciesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud ikepolicies delete unauthorized response has a 3xx status code
func (o *PcloudIkepoliciesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete unauthorized response has a 4xx status code
func (o *PcloudIkepoliciesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud ikepolicies delete unauthorized response has a 5xx status code
func (o *PcloudIkepoliciesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud ikepolicies delete unauthorized response a status code equal to that given
func (o *PcloudIkepoliciesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudIkepoliciesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudIkepoliciesDeleteUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudIkepoliciesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesDeleteForbidden creates a PcloudIkepoliciesDeleteForbidden with default headers values
func NewPcloudIkepoliciesDeleteForbidden() *PcloudIkepoliciesDeleteForbidden {
	return &PcloudIkepoliciesDeleteForbidden{}
}

/*
PcloudIkepoliciesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudIkepoliciesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud ikepolicies delete forbidden response has a 2xx status code
func (o *PcloudIkepoliciesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud ikepolicies delete forbidden response has a 3xx status code
func (o *PcloudIkepoliciesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete forbidden response has a 4xx status code
func (o *PcloudIkepoliciesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud ikepolicies delete forbidden response has a 5xx status code
func (o *PcloudIkepoliciesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud ikepolicies delete forbidden response a status code equal to that given
func (o *PcloudIkepoliciesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudIkepoliciesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudIkepoliciesDeleteForbidden) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteForbidden  %+v", 403, o.Payload)
}

func (o *PcloudIkepoliciesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesDeleteNotFound creates a PcloudIkepoliciesDeleteNotFound with default headers values
func NewPcloudIkepoliciesDeleteNotFound() *PcloudIkepoliciesDeleteNotFound {
	return &PcloudIkepoliciesDeleteNotFound{}
}

/*
PcloudIkepoliciesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudIkepoliciesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud ikepolicies delete not found response has a 2xx status code
func (o *PcloudIkepoliciesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud ikepolicies delete not found response has a 3xx status code
func (o *PcloudIkepoliciesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete not found response has a 4xx status code
func (o *PcloudIkepoliciesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud ikepolicies delete not found response has a 5xx status code
func (o *PcloudIkepoliciesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud ikepolicies delete not found response a status code equal to that given
func (o *PcloudIkepoliciesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudIkepoliciesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudIkepoliciesDeleteNotFound) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *PcloudIkepoliciesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesDeleteInternalServerError creates a PcloudIkepoliciesDeleteInternalServerError with default headers values
func NewPcloudIkepoliciesDeleteInternalServerError() *PcloudIkepoliciesDeleteInternalServerError {
	return &PcloudIkepoliciesDeleteInternalServerError{}
}

/*
PcloudIkepoliciesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudIkepoliciesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud ikepolicies delete internal server error response has a 2xx status code
func (o *PcloudIkepoliciesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud ikepolicies delete internal server error response has a 3xx status code
func (o *PcloudIkepoliciesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud ikepolicies delete internal server error response has a 4xx status code
func (o *PcloudIkepoliciesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud ikepolicies delete internal server error response has a 5xx status code
func (o *PcloudIkepoliciesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud ikepolicies delete internal server error response a status code equal to that given
func (o *PcloudIkepoliciesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudIkepoliciesDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudIkepoliciesDeleteInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesDeleteInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudIkepoliciesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
