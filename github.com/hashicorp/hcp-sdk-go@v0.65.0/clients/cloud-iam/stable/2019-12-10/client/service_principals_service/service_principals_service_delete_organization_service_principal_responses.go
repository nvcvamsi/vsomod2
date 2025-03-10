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

// ServicePrincipalsServiceDeleteOrganizationServicePrincipalReader is a Reader for the ServicePrincipalsServiceDeleteOrganizationServicePrincipal structure.
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceDeleteOrganizationServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalOK creates a ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK with default headers values
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalOK() *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK{}
}

/*
ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK struct {
	Payload models.HashicorpCloudIamDeleteOrganizationServicePrincipalResponse
}

// IsSuccess returns true when this service principals service delete organization service principal o k response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service delete organization service principal o k response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service delete organization service principal o k response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service delete organization service principal o k response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service delete organization service principal o k response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceDeleteOrganizationServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceDeleteOrganizationServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) GetPayload() models.HashicorpCloudIamDeleteOrganizationServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault creates a ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault with default headers values
func NewServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault(code int) *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault {
	return &ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service delete organization service principal default response
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service delete organization service principal default response has a 2xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service delete organization service principal default response has a 3xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service delete organization service principal default response has a 4xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service delete organization service principal default response has a 5xx status code
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service delete organization service principal default response a status code equal to that given
func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) Error() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_DeleteOrganizationServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) String() string {
	return fmt.Sprintf("[DELETE /iam/2019-12-10/organizations/{organization_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_DeleteOrganizationServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceDeleteOrganizationServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
