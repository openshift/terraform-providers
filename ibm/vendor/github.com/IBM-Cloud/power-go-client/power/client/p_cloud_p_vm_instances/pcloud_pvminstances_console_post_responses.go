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

// PcloudPvminstancesConsolePostReader is a Reader for the PcloudPvminstancesConsolePost structure.
type PcloudPvminstancesConsolePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesConsolePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPcloudPvminstancesConsolePostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPcloudPvminstancesConsolePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudPvminstancesConsolePostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPvminstancesConsolePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPvminstancesConsolePostCreated creates a PcloudPvminstancesConsolePostCreated with default headers values
func NewPcloudPvminstancesConsolePostCreated() *PcloudPvminstancesConsolePostCreated {
	return &PcloudPvminstancesConsolePostCreated{}
}

/*PcloudPvminstancesConsolePostCreated handles this case with default header values.

Created
*/
type PcloudPvminstancesConsolePostCreated struct {
	Payload *models.PVMInstanceConsole
}

func (o *PcloudPvminstancesConsolePostCreated) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostCreated  %+v", 201, o.Payload)
}

func (o *PcloudPvminstancesConsolePostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceConsole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostUnauthorized creates a PcloudPvminstancesConsolePostUnauthorized with default headers values
func NewPcloudPvminstancesConsolePostUnauthorized() *PcloudPvminstancesConsolePostUnauthorized {
	return &PcloudPvminstancesConsolePostUnauthorized{}
}

/*PcloudPvminstancesConsolePostUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudPvminstancesConsolePostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudPvminstancesConsolePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostUnprocessableEntity creates a PcloudPvminstancesConsolePostUnprocessableEntity with default headers values
func NewPcloudPvminstancesConsolePostUnprocessableEntity() *PcloudPvminstancesConsolePostUnprocessableEntity {
	return &PcloudPvminstancesConsolePostUnprocessableEntity{}
}

/*PcloudPvminstancesConsolePostUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudPvminstancesConsolePostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudPvminstancesConsolePostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesConsolePostInternalServerError creates a PcloudPvminstancesConsolePostInternalServerError with default headers values
func NewPcloudPvminstancesConsolePostInternalServerError() *PcloudPvminstancesConsolePostInternalServerError {
	return &PcloudPvminstancesConsolePostInternalServerError{}
}

/*PcloudPvminstancesConsolePostInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPvminstancesConsolePostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPvminstancesConsolePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/console][%d] pcloudPvminstancesConsolePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPvminstancesConsolePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
