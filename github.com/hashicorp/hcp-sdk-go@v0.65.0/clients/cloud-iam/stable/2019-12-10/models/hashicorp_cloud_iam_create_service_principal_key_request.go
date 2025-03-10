// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamCreateServicePrincipalKeyRequest CreateServicePrincipalKeyRequest is the request message used when creating a
// service principal key.
//
// swagger:model hashicorp.cloud.iam.CreateServicePrincipalKeyRequest
type HashicorpCloudIamCreateServicePrincipalKeyRequest struct {

	// parent_resource_name is the resource name of the service principal to
	// generate a key for.
	ParentResourceName string `json:"parent_resource_name,omitempty"`
}

// Validate validates this hashicorp cloud iam create service principal key request
func (m *HashicorpCloudIamCreateServicePrincipalKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam create service principal key request based on context it is used
func (m *HashicorpCloudIamCreateServicePrincipalKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalKeyRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateServicePrincipalKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
