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

// ServicePrincipalsServiceCreateProjectServicePrincipalKeyReader is a Reader for the ServicePrincipalsServiceCreateProjectServicePrincipalKey structure.
type ServicePrincipalsServiceCreateProjectServicePrincipalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceCreateProjectServicePrincipalKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceCreateProjectServicePrincipalKeyOK creates a ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK with default headers values
func NewServicePrincipalsServiceCreateProjectServicePrincipalKeyOK() *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK {
	return &ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK{}
}

/*
ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK struct {
	Payload *models.HashicorpCloudIamCreateProjectServicePrincipalKeyResponse
}

// IsSuccess returns true when this service principals service create project service principal key o k response has a 2xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service create project service principal key o k response has a 3xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service create project service principal key o k response has a 4xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service create project service principal key o k response has a 5xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service create project service principal key o k response a status code equal to that given
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys][%d] servicePrincipalsServiceCreateProjectServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys][%d] servicePrincipalsServiceCreateProjectServicePrincipalKeyOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) GetPayload() *models.HashicorpCloudIamCreateProjectServicePrincipalKeyResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamCreateProjectServicePrincipalKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault creates a ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault with default headers values
func NewServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault(code int) *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault {
	return &ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service create project service principal key default response
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service create project service principal key default response has a 2xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service create project service principal key default response has a 3xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service create project service principal key default response has a 4xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service create project service principal key default response has a 5xx status code
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service create project service principal key default response a status code equal to that given
func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys][%d] ServicePrincipalsService_CreateProjectServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principal-keys][%d] ServicePrincipalsService_CreateProjectServicePrincipalKey default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceCreateProjectServicePrincipalKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
