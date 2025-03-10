// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceCreateServicePrincipalKey2Reader is a Reader for the ServicePrincipalsServiceCreateServicePrincipalKey2 structure.
type ServicePrincipalsServiceCreateServicePrincipalKey2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceCreateServicePrincipalKey2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceCreateServicePrincipalKey2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceCreateServicePrincipalKey2OK creates a ServicePrincipalsServiceCreateServicePrincipalKey2OK with default headers values
func NewServicePrincipalsServiceCreateServicePrincipalKey2OK() *ServicePrincipalsServiceCreateServicePrincipalKey2OK {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2OK{}
}

/*
ServicePrincipalsServiceCreateServicePrincipalKey2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceCreateServicePrincipalKey2OK struct {
	Payload *models.HashicorpCloudIamCreateServicePrincipalKeyResponse
}

// IsSuccess returns true when this service principals service create service principal key2 o k response has a 2xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service create service principal key2 o k response has a 3xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service create service principal key2 o k response has a 4xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service create service principal key2 o k response has a 5xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service create service principal key2 o k response a status code equal to that given
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/keys][%d] servicePrincipalsServiceCreateServicePrincipalKey2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/keys][%d] servicePrincipalsServiceCreateServicePrincipalKey2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) GetPayload() *models.HashicorpCloudIamCreateServicePrincipalKeyResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateServicePrincipalKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceCreateServicePrincipalKey2Default creates a ServicePrincipalsServiceCreateServicePrincipalKey2Default with default headers values
func NewServicePrincipalsServiceCreateServicePrincipalKey2Default(code int) *ServicePrincipalsServiceCreateServicePrincipalKey2Default {
	return &ServicePrincipalsServiceCreateServicePrincipalKey2Default{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceCreateServicePrincipalKey2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceCreateServicePrincipalKey2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service create service principal key2 default response
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service create service principal key2 default response has a 2xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service create service principal key2 default response has a 3xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service create service principal key2 default response has a 4xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service create service principal key2 default response has a 5xx status code
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service create service principal key2 default response a status code equal to that given
func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) Error() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/keys][%d] ServicePrincipalsService_CreateServicePrincipalKey2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) String() string {
	return fmt.Sprintf("[POST /2019-12-10/{parent_resource_name}/keys][%d] ServicePrincipalsService_CreateServicePrincipalKey2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateServicePrincipalKey2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
