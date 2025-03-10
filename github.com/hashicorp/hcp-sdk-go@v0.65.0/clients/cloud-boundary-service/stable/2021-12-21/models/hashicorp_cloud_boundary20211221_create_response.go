// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudBoundary20211221CreateResponse CreateResponse is a response on cluster creation.
//
// swagger:model hashicorp.cloud.boundary_20211221.CreateResponse
type HashicorpCloudBoundary20211221CreateResponse struct {

	// cluster_id is the ID of the Boundary cluster that users set on creation.
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster_url is the generated vanity url.
	ClusterURL string `json:"cluster_url,omitempty"`

	// operation tracks the asynchronous creation of the Boundary cluster.
	Operation *cloud.HashicorpCloudOperationOperation `json:"operation,omitempty"`
}

// Validate validates this hashicorp cloud boundary 20211221 create response
func (m *HashicorpCloudBoundary20211221CreateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudBoundary20211221CreateResponse) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud boundary 20211221 create response based on the context it is used
func (m *HashicorpCloudBoundary20211221CreateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudBoundary20211221CreateResponse) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221CreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221CreateResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudBoundary20211221CreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
