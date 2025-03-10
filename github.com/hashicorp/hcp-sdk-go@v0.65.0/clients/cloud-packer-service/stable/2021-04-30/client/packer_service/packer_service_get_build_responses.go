// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceGetBuildReader is a Reader for the PackerServiceGetBuild structure.
type PackerServiceGetBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceGetBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceGetBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceGetBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceGetBuildOK creates a PackerServiceGetBuildOK with default headers values
func NewPackerServiceGetBuildOK() *PackerServiceGetBuildOK {
	return &PackerServiceGetBuildOK{}
}

/*
PackerServiceGetBuildOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceGetBuildOK struct {
	Payload *models.HashicorpCloudPackerGetBuildResponse
}

// IsSuccess returns true when this packer service get build o k response has a 2xx status code
func (o *PackerServiceGetBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service get build o k response has a 3xx status code
func (o *PackerServiceGetBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service get build o k response has a 4xx status code
func (o *PackerServiceGetBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service get build o k response has a 5xx status code
func (o *PackerServiceGetBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service get build o k response a status code equal to that given
func (o *PackerServiceGetBuildOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceGetBuildOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] packerServiceGetBuildOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetBuildOK) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] packerServiceGetBuildOK  %+v", 200, o.Payload)
}

func (o *PackerServiceGetBuildOK) GetPayload() *models.HashicorpCloudPackerGetBuildResponse {
	return o.Payload
}

func (o *PackerServiceGetBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerGetBuildResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceGetBuildDefault creates a PackerServiceGetBuildDefault with default headers values
func NewPackerServiceGetBuildDefault(code int) *PackerServiceGetBuildDefault {
	return &PackerServiceGetBuildDefault{
		_statusCode: code,
	}
}

/*
PackerServiceGetBuildDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceGetBuildDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service get build default response
func (o *PackerServiceGetBuildDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service get build default response has a 2xx status code
func (o *PackerServiceGetBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service get build default response has a 3xx status code
func (o *PackerServiceGetBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service get build default response has a 4xx status code
func (o *PackerServiceGetBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service get build default response has a 5xx status code
func (o *PackerServiceGetBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service get build default response a status code equal to that given
func (o *PackerServiceGetBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceGetBuildDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] PackerService_GetBuild default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetBuildDefault) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] PackerService_GetBuild default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceGetBuildDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceGetBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
