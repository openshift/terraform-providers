// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_onboarding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVolumeOnboardingGetallReader is a Reader for the PcloudVolumeOnboardingGetall structure.
type PcloudVolumeOnboardingGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVolumeOnboardingGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudVolumeOnboardingGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVolumeOnboardingGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVolumeOnboardingGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVolumeOnboardingGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVolumeOnboardingGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVolumeOnboardingGetallOK creates a PcloudVolumeOnboardingGetallOK with default headers values
func NewPcloudVolumeOnboardingGetallOK() *PcloudVolumeOnboardingGetallOK {
	return &PcloudVolumeOnboardingGetallOK{}
}

/*
PcloudVolumeOnboardingGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudVolumeOnboardingGetallOK struct {
	Payload *models.VolumeOnboardings
}

// IsSuccess returns true when this pcloud volume onboarding getall o k response has a 2xx status code
func (o *PcloudVolumeOnboardingGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud volume onboarding getall o k response has a 3xx status code
func (o *PcloudVolumeOnboardingGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volume onboarding getall o k response has a 4xx status code
func (o *PcloudVolumeOnboardingGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volume onboarding getall o k response has a 5xx status code
func (o *PcloudVolumeOnboardingGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volume onboarding getall o k response a status code equal to that given
func (o *PcloudVolumeOnboardingGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudVolumeOnboardingGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallOK) GetPayload() *models.VolumeOnboardings {
	return o.Payload
}

func (o *PcloudVolumeOnboardingGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeOnboardings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingGetallBadRequest creates a PcloudVolumeOnboardingGetallBadRequest with default headers values
func NewPcloudVolumeOnboardingGetallBadRequest() *PcloudVolumeOnboardingGetallBadRequest {
	return &PcloudVolumeOnboardingGetallBadRequest{}
}

/*
PcloudVolumeOnboardingGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVolumeOnboardingGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volume onboarding getall bad request response has a 2xx status code
func (o *PcloudVolumeOnboardingGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volume onboarding getall bad request response has a 3xx status code
func (o *PcloudVolumeOnboardingGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volume onboarding getall bad request response has a 4xx status code
func (o *PcloudVolumeOnboardingGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volume onboarding getall bad request response has a 5xx status code
func (o *PcloudVolumeOnboardingGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volume onboarding getall bad request response a status code equal to that given
func (o *PcloudVolumeOnboardingGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudVolumeOnboardingGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingGetallUnauthorized creates a PcloudVolumeOnboardingGetallUnauthorized with default headers values
func NewPcloudVolumeOnboardingGetallUnauthorized() *PcloudVolumeOnboardingGetallUnauthorized {
	return &PcloudVolumeOnboardingGetallUnauthorized{}
}

/*
PcloudVolumeOnboardingGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVolumeOnboardingGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volume onboarding getall unauthorized response has a 2xx status code
func (o *PcloudVolumeOnboardingGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volume onboarding getall unauthorized response has a 3xx status code
func (o *PcloudVolumeOnboardingGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volume onboarding getall unauthorized response has a 4xx status code
func (o *PcloudVolumeOnboardingGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volume onboarding getall unauthorized response has a 5xx status code
func (o *PcloudVolumeOnboardingGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volume onboarding getall unauthorized response a status code equal to that given
func (o *PcloudVolumeOnboardingGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudVolumeOnboardingGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingGetallForbidden creates a PcloudVolumeOnboardingGetallForbidden with default headers values
func NewPcloudVolumeOnboardingGetallForbidden() *PcloudVolumeOnboardingGetallForbidden {
	return &PcloudVolumeOnboardingGetallForbidden{}
}

/*
PcloudVolumeOnboardingGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVolumeOnboardingGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volume onboarding getall forbidden response has a 2xx status code
func (o *PcloudVolumeOnboardingGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volume onboarding getall forbidden response has a 3xx status code
func (o *PcloudVolumeOnboardingGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volume onboarding getall forbidden response has a 4xx status code
func (o *PcloudVolumeOnboardingGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud volume onboarding getall forbidden response has a 5xx status code
func (o *PcloudVolumeOnboardingGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud volume onboarding getall forbidden response a status code equal to that given
func (o *PcloudVolumeOnboardingGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudVolumeOnboardingGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVolumeOnboardingGetallInternalServerError creates a PcloudVolumeOnboardingGetallInternalServerError with default headers values
func NewPcloudVolumeOnboardingGetallInternalServerError() *PcloudVolumeOnboardingGetallInternalServerError {
	return &PcloudVolumeOnboardingGetallInternalServerError{}
}

/*
PcloudVolumeOnboardingGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVolumeOnboardingGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud volume onboarding getall internal server error response has a 2xx status code
func (o *PcloudVolumeOnboardingGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud volume onboarding getall internal server error response has a 3xx status code
func (o *PcloudVolumeOnboardingGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud volume onboarding getall internal server error response has a 4xx status code
func (o *PcloudVolumeOnboardingGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud volume onboarding getall internal server error response has a 5xx status code
func (o *PcloudVolumeOnboardingGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud volume onboarding getall internal server error response a status code equal to that given
func (o *PcloudVolumeOnboardingGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudVolumeOnboardingGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/onboarding][%d] pcloudVolumeOnboardingGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVolumeOnboardingGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVolumeOnboardingGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
