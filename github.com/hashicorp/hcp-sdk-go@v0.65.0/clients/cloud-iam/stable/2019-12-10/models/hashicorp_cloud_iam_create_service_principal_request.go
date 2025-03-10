// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamCreateServicePrincipalRequest CreateServicePrincipalRequest is the request message used when creating a
// service principal.
//
// swagger:model hashicorp.cloud.iam.CreateServicePrincipalRequest
type HashicorpCloudIamCreateServicePrincipalRequest struct {

	// name is the customer-chosen name for this service principal.
	Name string `json:"name,omitempty"`

	// parent_resource_name is the resource name of the project or organization
	// at which the service principal should be created (e.g.
	// "project/<project_id>" or "organization/<organization_id>")
	ParentResourceName string `json:"parent_resource_name,omitempty"`
}

// Validate validates this hashicorp cloud iam create service principal request
func (m *HashicorpCloudIamCreateServicePrincipalRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam create service principal request based on context it is used
func (m *HashicorpCloudIamCreateServicePrincipalRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateServicePrincipalRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateServicePrincipalRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
