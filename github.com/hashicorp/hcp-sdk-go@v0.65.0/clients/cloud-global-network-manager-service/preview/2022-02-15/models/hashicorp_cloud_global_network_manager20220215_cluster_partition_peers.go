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

// HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers hashicorp cloud global network manager 20220215 cluster partition peers
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ClusterPartitionPeers
type HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers struct {

	// cluster_id is the user settable GNM cluster name.
	ClusterID string `json:"cluster_id,omitempty"`

	// licensing is the Consul licensing information
	Licensing *HashicorpCloudGlobalNetworkManager20220215Licensing `json:"licensing,omitempty"`

	// partition_name is the name of the admin partition on the cluster
	PartitionName string `json:"partition_name,omitempty"`

	// peers is the list of cluster peers connected to the above cluster and partition
	Peers []*HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeer `json:"peers"`
}

// Validate validates this hashicorp cloud global network manager 20220215 cluster partition peers
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicensing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) validateLicensing(formats strfmt.Registry) error {
	if swag.IsZero(m.Licensing) { // not required
		return nil
	}

	if m.Licensing != nil {
		if err := m.Licensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) validatePeers(formats strfmt.Registry) error {
	if swag.IsZero(m.Peers) { // not required
		return nil
	}

	for i := 0; i < len(m.Peers); i++ {
		if swag.IsZero(m.Peers[i]) { // not required
			continue
		}

		if m.Peers[i] != nil {
			if err := m.Peers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 cluster partition peers based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLicensing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) contextValidateLicensing(ctx context.Context, formats strfmt.Registry) error {

	if m.Licensing != nil {
		if err := m.Licensing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) contextValidatePeers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Peers); i++ {

		if m.Peers[i] != nil {
			if err := m.Peers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("peers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("peers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ClusterPartitionPeers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
