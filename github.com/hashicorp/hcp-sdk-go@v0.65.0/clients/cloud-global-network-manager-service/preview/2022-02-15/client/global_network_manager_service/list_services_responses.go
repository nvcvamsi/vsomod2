// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListServicesReader is a Reader for the ListServices structure.
type ListServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListServicesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListServicesOK creates a ListServicesOK with default headers values
func NewListServicesOK() *ListServicesOK {
	return &ListServicesOK{}
}

/*
ListServicesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListServicesOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215ListServicesResponse
}

// IsSuccess returns true when this list services o k response has a 2xx status code
func (o *ListServicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list services o k response has a 3xx status code
func (o *ListServicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list services o k response has a 4xx status code
func (o *ListServicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list services o k response has a 5xx status code
func (o *ListServicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list services o k response a status code equal to that given
func (o *ListServicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListServicesOK) Error() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{project_resource_name}/services][%d] listServicesOK  %+v", 200, o.Payload)
}

func (o *ListServicesOK) String() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{project_resource_name}/services][%d] listServicesOK  %+v", 200, o.Payload)
}

func (o *ListServicesOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215ListServicesResponse {
	return o.Payload
}

func (o *ListServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215ListServicesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListServicesDefault creates a ListServicesDefault with default headers values
func NewListServicesDefault(code int) *ListServicesDefault {
	return &ListServicesDefault{
		_statusCode: code,
	}
}

/*
ListServicesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListServicesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the list services default response
func (o *ListServicesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list services default response has a 2xx status code
func (o *ListServicesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list services default response has a 3xx status code
func (o *ListServicesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list services default response has a 4xx status code
func (o *ListServicesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list services default response has a 5xx status code
func (o *ListServicesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list services default response a status code equal to that given
func (o *ListServicesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListServicesDefault) Error() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{project_resource_name}/services][%d] ListServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListServicesDefault) String() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{project_resource_name}/services][%d] ListServices default  %+v", o._statusCode, o.Payload)
}

func (o *ListServicesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ListServicesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
