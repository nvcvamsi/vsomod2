// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudResourcemanagerOrganizationOwner OrganizationOwner identifies the user who owns the organization.
//
// swagger:model hashicorp.cloud.resourcemanager.OrganizationOwner
type HashicorpCloudResourcemanagerOrganizationOwner struct {

	// User is the email of the user who is the owner.
	User string `json:"user,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager organization owner
func (m *HashicorpCloudResourcemanagerOrganizationOwner) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud resourcemanager organization owner based on context it is used
func (m *HashicorpCloudResourcemanagerOrganizationOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerOrganizationOwner) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerOrganizationOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
