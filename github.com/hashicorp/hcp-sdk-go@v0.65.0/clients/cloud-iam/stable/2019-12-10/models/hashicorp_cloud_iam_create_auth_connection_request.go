// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudIamCreateAuthConnectionRequest CreateAuthConnectionRequest creates an AuthConnection.
//
// swagger:model hashicorp.cloud.iam.CreateAuthConnectionRequest
type HashicorpCloudIamCreateAuthConnectionRequest struct {

	// client_id is the ID of the client for the connection.
	ClientID string `json:"client_id,omitempty"`

	// client_secret is the secret for the client above.
	ClientSecret string `json:"client_secret,omitempty"`

	// email_domain associated with the of the connection.
	EmailDomain string `json:"email_domain,omitempty"`

	// issuer is the URL of the identity provider.
	Issuer string `json:"issuer,omitempty"`

	// organization_id is the organization that will own the authentication connection.
	OrganizationID string `json:"organization_id,omitempty"`
}

// Validate validates this hashicorp cloud iam create auth connection request
func (m *HashicorpCloudIamCreateAuthConnectionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud iam create auth connection request based on context it is used
func (m *HashicorpCloudIamCreateAuthConnectionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateAuthConnectionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateAuthConnectionRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateAuthConnectionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
