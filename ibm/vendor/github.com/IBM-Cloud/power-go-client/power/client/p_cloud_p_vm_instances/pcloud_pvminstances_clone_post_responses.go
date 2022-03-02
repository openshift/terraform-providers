// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesClonePostReader is a Reader for the PcloudPvminstancesClonePost structure.
type PcloudPvminstancesClonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesClonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPcloudPvminstancesClonePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPvminstancesClonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudPvminstancesClonePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPcloudPvminstancesClonePostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudPvminstancesClonePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesClonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesClonePostAccepted creates a PcloudPvminstancesClonePostAccepted with default headers values
func NewPcloudPvminstancesClonePostAccepted() *PcloudPvminstancesClonePostAccepted {
	return &PcloudPvminstancesClonePostAccepted{}
}

/*PcloudPvminstancesClonePostAccepted handles this case with default header values.

Accepted
*/
type PcloudPvminstancesClonePostAccepted struct {
	Payload *models.PVMInstance
}

func (o *PcloudPvminstancesClonePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudPvminstancesClonePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostBadRequest creates a PcloudPvminstancesClonePostBadRequest with default headers values
func NewPcloudPvminstancesClonePostBadRequest() *PcloudPvminstancesClonePostBadRequest {
	return &PcloudPvminstancesClonePostBadRequest{}
}

/*PcloudPvminstancesClonePostBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPvminstancesClonePostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesClonePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPvminstancesClonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostUnauthorized creates a PcloudPvminstancesClonePostUnauthorized with default headers values
func NewPcloudPvminstancesClonePostUnauthorized() *PcloudPvminstancesClonePostUnauthorized {
	return &PcloudPvminstancesClonePostUnauthorized{}
}

/*PcloudPvminstancesClonePostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudPvminstancesClonePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesClonePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostConflict creates a PcloudPvminstancesClonePostConflict with default headers values
func NewPcloudPvminstancesClonePostConflict() *PcloudPvminstancesClonePostConflict {
	return &PcloudPvminstancesClonePostConflict{}
}

/*PcloudPvminstancesClonePostConflict handles this case with default header values.

Conflict
*/
type PcloudPvminstancesClonePostConflict struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesClonePostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostConflict  %+v", 409, o.Payload)
}

func (o *PcloudPvminstancesClonePostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostUnprocessableEntity creates a PcloudPvminstancesClonePostUnprocessableEntity with default headers values
func NewPcloudPvminstancesClonePostUnprocessableEntity() *PcloudPvminstancesClonePostUnprocessableEntity {
	return &PcloudPvminstancesClonePostUnprocessableEntity{}
}

/*PcloudPvminstancesClonePostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesClonePostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesClonePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesClonePostInternalServerError creates a PcloudPvminstancesClonePostInternalServerError with default headers values
func NewPcloudPvminstancesClonePostInternalServerError() *PcloudPvminstancesClonePostInternalServerError {
	return &PcloudPvminstancesClonePostInternalServerError{}
}

/*PcloudPvminstancesClonePostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesClonePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesClonePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/clone][%d] pcloudPvminstancesClonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesClonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
