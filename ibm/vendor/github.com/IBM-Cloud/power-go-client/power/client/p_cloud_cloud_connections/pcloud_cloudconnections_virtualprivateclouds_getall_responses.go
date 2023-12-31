// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsVirtualprivatecloudsGetallReader is a Reader for the PcloudCloudconnectionsVirtualprivatecloudsGetall structure.
type PcloudCloudconnectionsVirtualprivatecloudsGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallOK creates a PcloudCloudconnectionsVirtualprivatecloudsGetallOK with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallOK() *PcloudCloudconnectionsVirtualprivatecloudsGetallOK {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallOK{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallOK struct {
	Payload *models.CloudConnectionVirtualPrivateClouds
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall o k response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall o k response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall o k response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall o k response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall o k response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) GetPayload() *models.CloudConnectionVirtualPrivateClouds {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnectionVirtualPrivateClouds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest creates a PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest() *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall bad request response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall bad request response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall bad request response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall bad request response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall bad request response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized creates a PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized() *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallForbidden creates a PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallForbidden() *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall forbidden response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall forbidden response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall forbidden response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall forbidden response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall forbidden response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout creates a PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout() *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout  %+v", 408, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError creates a PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError() *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout creates a PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout with default headers values
func NewPcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout() *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout {
	return &PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout{}
}

/*
PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout
*/
type PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections virtualprivateclouds getall gateway timeout response has a 2xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections virtualprivateclouds getall gateway timeout response has a 3xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections virtualprivateclouds getall gateway timeout response has a 4xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections virtualprivateclouds getall gateway timeout response has a 5xx status code
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections virtualprivateclouds getall gateway timeout response a status code equal to that given
func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections-virtual-private-clouds][%d] pcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsVirtualprivatecloudsGetallGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
