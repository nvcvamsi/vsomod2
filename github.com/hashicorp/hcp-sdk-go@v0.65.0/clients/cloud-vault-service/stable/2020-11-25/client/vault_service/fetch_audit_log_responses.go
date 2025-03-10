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

// FetchAuditLogReader is a Reader for the FetchAuditLog structure.
type FetchAuditLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchAuditLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFetchAuditLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFetchAuditLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFetchAuditLogOK creates a FetchAuditLogOK with default headers values
func NewFetchAuditLogOK() *FetchAuditLogOK {
	return &FetchAuditLogOK{}
}

/*
FetchAuditLogOK describes a response with status code 200, with default header values.

A successful response.
*/
type FetchAuditLogOK struct {
	Payload *models.HashicorpCloudVault20201125FetchAuditLogResponse
}

// IsSuccess returns true when this fetch audit log o k response has a 2xx status code
func (o *FetchAuditLogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fetch audit log o k response has a 3xx status code
func (o *FetchAuditLogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch audit log o k response has a 4xx status code
func (o *FetchAuditLogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fetch audit log o k response has a 5xx status code
func (o *FetchAuditLogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch audit log o k response a status code equal to that given
func (o *FetchAuditLogOK) IsCode(code int) bool {
	return code == 200
}

func (o *FetchAuditLogOK) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog][%d] fetchAuditLogOK  %+v", 200, o.Payload)
}

func (o *FetchAuditLogOK) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog][%d] fetchAuditLogOK  %+v", 200, o.Payload)
}

func (o *FetchAuditLogOK) GetPayload() *models.HashicorpCloudVault20201125FetchAuditLogResponse {
	return o.Payload
}

func (o *FetchAuditLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVault20201125FetchAuditLogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchAuditLogDefault creates a FetchAuditLogDefault with default headers values
func NewFetchAuditLogDefault(code int) *FetchAuditLogDefault {
	return &FetchAuditLogDefault{
		_statusCode: code,
	}
}

/*
FetchAuditLogDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type FetchAuditLogDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the fetch audit log default response
func (o *FetchAuditLogDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fetch audit log default response has a 2xx status code
func (o *FetchAuditLogDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fetch audit log default response has a 3xx status code
func (o *FetchAuditLogDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fetch audit log default response has a 4xx status code
func (o *FetchAuditLogDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fetch audit log default response has a 5xx status code
func (o *FetchAuditLogDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fetch audit log default response a status code equal to that given
func (o *FetchAuditLogDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FetchAuditLogDefault) Error() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog][%d] FetchAuditLog default  %+v", o._statusCode, o.Payload)
}

func (o *FetchAuditLogDefault) String() string {
	return fmt.Sprintf("[POST /vault/2020-11-25/organizations/{location.organization_id}/projects/{location.project_id}/clusters/{cluster_id}/auditlog][%d] FetchAuditLog default  %+v", o._statusCode, o.Payload)
}

func (o *FetchAuditLogDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *FetchAuditLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
