// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter hashicorp cloud global network manager 20220215 get service summaries request filter
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.GetServiceSummariesRequest.Filter
type HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter struct {

	// clusters matches summaries on the cluster name
	Clusters []string `json:"clusters"`

	// kinds matches service kind. If empty, this defaults to all kinds
	Kinds []*HashicorpCloudGlobalNetworkManager20220215ServiceSummaryKind `json:"kinds"`

	// name_substr matches summaries that contain a case-insensitive substring in their service name
	NameSubstr string `json:"name_substr,omitempty"`

	// namespaces matches summaries on the namespace name
	Namespaces []string `json:"namespaces"`

	// partitions matches summaries on the partition name
	Partitions []string `json:"partitions"`
}

// Validate validates this hashicorp cloud global network manager 20220215 get service summaries request filter
func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKinds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) validateKinds(formats strfmt.Registry) error {
	if swag.IsZero(m.Kinds) { // not required
		return nil
	}

	for i := 0; i < len(m.Kinds); i++ {
		if swag.IsZero(m.Kinds[i]) { // not required
			continue
		}

		if m.Kinds[i] != nil {
			if err := m.Kinds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 get service summaries request filter based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) contextValidateKinds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Kinds); i++ {

		if m.Kinds[i] != nil {
			if err := m.Kinds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("kinds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("kinds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215GetServiceSummariesRequestFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
