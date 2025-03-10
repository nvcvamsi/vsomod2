// Code generated by go-swagger; DO NOT EDIT.

package organization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// OrganizationServiceTestIamPermissionsReader is a Reader for the OrganizationServiceTestIamPermissions structure.
type OrganizationServiceTestIamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrganizationServiceTestIamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrganizationServiceTestIamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOrganizationServiceTestIamPermissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOrganizationServiceTestIamPermissionsOK creates a OrganizationServiceTestIamPermissionsOK with default headers values
func NewOrganizationServiceTestIamPermissionsOK() *OrganizationServiceTestIamPermissionsOK {
	return &OrganizationServiceTestIamPermissionsOK{}
}

/*
OrganizationServiceTestIamPermissionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type OrganizationServiceTestIamPermissionsOK struct {
	Payload *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse
}

// IsSuccess returns true when this organization service test iam permissions o k response has a 2xx status code
func (o *OrganizationServiceTestIamPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this organization service test iam permissions o k response has a 3xx status code
func (o *OrganizationServiceTestIamPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this organization service test iam permissions o k response has a 4xx status code
func (o *OrganizationServiceTestIamPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this organization service test iam permissions o k response has a 5xx status code
func (o *OrganizationServiceTestIamPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this organization service test iam permissions o k response a status code equal to that given
func (o *OrganizationServiceTestIamPermissionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *OrganizationServiceTestIamPermissionsOK) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] organizationServiceTestIamPermissionsOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsOK) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] organizationServiceTestIamPermissionsOK  %+v", 200, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsOK) GetPayload() *models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse {
	return o.Payload
}

func (o *OrganizationServiceTestIamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudResourcemanagerOrganizationTestIamPermissionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrganizationServiceTestIamPermissionsDefault creates a OrganizationServiceTestIamPermissionsDefault with default headers values
func NewOrganizationServiceTestIamPermissionsDefault(code int) *OrganizationServiceTestIamPermissionsDefault {
	return &OrganizationServiceTestIamPermissionsDefault{
		_statusCode: code,
	}
}

/*
OrganizationServiceTestIamPermissionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OrganizationServiceTestIamPermissionsDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the organization service test iam permissions default response
func (o *OrganizationServiceTestIamPermissionsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this organization service test iam permissions default response has a 2xx status code
func (o *OrganizationServiceTestIamPermissionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this organization service test iam permissions default response has a 3xx status code
func (o *OrganizationServiceTestIamPermissionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this organization service test iam permissions default response has a 4xx status code
func (o *OrganizationServiceTestIamPermissionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this organization service test iam permissions default response has a 5xx status code
func (o *OrganizationServiceTestIamPermissionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this organization service test iam permissions default response a status code equal to that given
func (o *OrganizationServiceTestIamPermissionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OrganizationServiceTestIamPermissionsDefault) Error() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] OrganizationService_TestIamPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsDefault) String() string {
	return fmt.Sprintf("[POST /resource-manager/2019-12-10/organizations/{id}/test-iam-permissions][%d] OrganizationService_TestIamPermissions default  %+v", o._statusCode, o.Payload)
}

func (o *OrganizationServiceTestIamPermissionsDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *OrganizationServiceTestIamPermissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
OrganizationServiceTestIamPermissionsBody see OrganizationService.TestIamPermissions
swagger:model OrganizationServiceTestIamPermissionsBody
*/
type OrganizationServiceTestIamPermissionsBody struct {

	// Permissions to test.
	Permissions []string `json:"permissions"`
}

// Validate validates this organization service test iam permissions body
func (o *OrganizationServiceTestIamPermissionsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this organization service test iam permissions body based on context it is used
func (o *OrganizationServiceTestIamPermissionsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *OrganizationServiceTestIamPermissionsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *OrganizationServiceTestIamPermissionsBody) UnmarshalBinary(b []byte) error {
	var res OrganizationServiceTestIamPermissionsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
