// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// GetSyncInstallationReader is a Reader for the GetSyncInstallation structure.
type GetSyncInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSyncInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSyncInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSyncInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSyncInstallationOK creates a GetSyncInstallationOK with default headers values
func NewGetSyncInstallationOK() *GetSyncInstallationOK {
	return &GetSyncInstallationOK{}
}

/*
GetSyncInstallationOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSyncInstallationOK struct {
	Payload *models.Secrets20230613GetSyncInstallationResponse
}

// IsSuccess returns true when this get sync installation o k response has a 2xx status code
func (o *GetSyncInstallationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sync installation o k response has a 3xx status code
func (o *GetSyncInstallationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sync installation o k response has a 4xx status code
func (o *GetSyncInstallationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sync installation o k response has a 5xx status code
func (o *GetSyncInstallationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sync installation o k response a status code equal to that given
func (o *GetSyncInstallationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSyncInstallationOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations/{name}][%d] getSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *GetSyncInstallationOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations/{name}][%d] getSyncInstallationOK  %+v", 200, o.Payload)
}

func (o *GetSyncInstallationOK) GetPayload() *models.Secrets20230613GetSyncInstallationResponse {
	return o.Payload
}

func (o *GetSyncInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613GetSyncInstallationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSyncInstallationDefault creates a GetSyncInstallationDefault with default headers values
func NewGetSyncInstallationDefault(code int) *GetSyncInstallationDefault {
	return &GetSyncInstallationDefault{
		_statusCode: code,
	}
}

/*
GetSyncInstallationDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetSyncInstallationDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the get sync installation default response
func (o *GetSyncInstallationDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get sync installation default response has a 2xx status code
func (o *GetSyncInstallationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get sync installation default response has a 3xx status code
func (o *GetSyncInstallationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get sync installation default response has a 4xx status code
func (o *GetSyncInstallationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get sync installation default response has a 5xx status code
func (o *GetSyncInstallationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get sync installation default response a status code equal to that given
func (o *GetSyncInstallationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetSyncInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations/{name}][%d] GetSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GetSyncInstallationDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/sync/installations/{name}][%d] GetSyncInstallation default  %+v", o._statusCode, o.Payload)
}

func (o *GetSyncInstallationDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *GetSyncInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
