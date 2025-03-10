// Code generated by go-swagger; DO NOT EDIT.

package vault_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-service/stable/2020-11-25/models"
)

// GetAvailableTemplatesReader is a Reader for the GetAvailableTemplates structure.
type GetAvailableTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAvailableTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAvailableTemplatesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAvailableTemplatesOK creates a GetAvailableTemplatesOK with default headers values
func NewGetAvailableTemplatesOK() *GetAvailableTemplatesOK {
	return &GetAvailableTemplatesOK{}
}

/*
GetAvailableTemplatesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetAvailableTemplatesOK struct {
	Payload *models.HashicorpCloudVault20201125GetAvailableTemplatesResponse
}

// IsSuccess returns true when this get available templates o k response has a 2xx status code
func (o *GetAvailableTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get available templates o k response has a 3xx status code
func (o *GetAvailableTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get available templates o k response has a 4xx status code
func (o *GetAvailableTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get available templates o k response has a 5xx status code
func (o *GetAvailableTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get available templates o k response a status code equal to that given
func (o *GetAvailableTemplatesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAvailableTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/templates][%d] getAvailableTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAvailableTemplatesOK) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/templates][%d] getAvailableTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetAvailableTemplatesOK) GetPayload() *models.HashicorpCloudVault20201125GetAvailableTemplatesResponse {
	return o.Payload
}

func (o *GetAvailableTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125GetAvailableTemplatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAvailableTemplatesDefault creates a GetAvailableTemplatesDefault with default headers values
func NewGetAvailableTemplatesDefault(code int) *GetAvailableTemplatesDefault {
	return &GetAvailableTemplatesDefault{
		_statusCode: code,
	}
}

/*
GetAvailableTemplatesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetAvailableTemplatesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the get available templates default response
func (o *GetAvailableTemplatesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get available templates default response has a 2xx status code
func (o *GetAvailableTemplatesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get available templates default response has a 3xx status code
func (o *GetAvailableTemplatesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get available templates default response has a 4xx status code
func (o *GetAvailableTemplatesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get available templates default response has a 5xx status code
func (o *GetAvailableTemplatesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get available templates default response a status code equal to that given
func (o *GetAvailableTemplatesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetAvailableTemplatesDefault) Error() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/templates][%d] GetAvailableTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailableTemplatesDefault) String() string {
	return fmt.Sprintf("[GET /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/templates][%d] GetAvailableTemplates default  %+v", o._statusCode, o.Payload)
}

func (o *GetAvailableTemplatesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *GetAvailableTemplatesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
