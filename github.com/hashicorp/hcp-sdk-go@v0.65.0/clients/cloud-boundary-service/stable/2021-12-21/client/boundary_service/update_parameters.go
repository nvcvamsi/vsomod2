// Code generated by go-swagger; DO NOT EDIT.

package boundary_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-boundary-service/stable/2021-12-21/models"
)

// NewUpdateParams creates a new UpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateParams() *UpdateParams {
	return &UpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateParamsWithTimeout creates a new UpdateParams object
// with the ability to set a timeout on a request.
func NewUpdateParamsWithTimeout(timeout time.Duration) *UpdateParams {
	return &UpdateParams{
		timeout: timeout,
	}
}

// NewUpdateParamsWithContext creates a new UpdateParams object
// with the ability to set a context for a request.
func NewUpdateParamsWithContext(ctx context.Context) *UpdateParams {
	return &UpdateParams{
		Context: ctx,
	}
}

// NewUpdateParamsWithHTTPClient creates a new UpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateParamsWithHTTPClient(client *http.Client) *UpdateParams {
	return &UpdateParams{
		HTTPClient: client,
	}
}

/*
UpdateParams contains all the parameters to send to the API endpoint

	for the update operation.

	Typically these are written to a http.Request.
*/
type UpdateParams struct {

	// Body.
	Body *models.HashicorpCloudBoundary20211221UpdateRequest

	/* ClusterID.

	   cluster_id is the id of the cluster set by user on creation.
	*/
	ClusterID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParams) WithDefaults() *UpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update params
func (o *UpdateParams) WithTimeout(timeout time.Duration) *UpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update params
func (o *UpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update params
func (o *UpdateParams) WithContext(ctx context.Context) *UpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update params
func (o *UpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) WithHTTPClient(client *http.Client) *UpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update params
func (o *UpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update params
func (o *UpdateParams) WithBody(body *models.HashicorpCloudBoundary20211221UpdateRequest) *UpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update params
func (o *UpdateParams) SetBody(body *models.HashicorpCloudBoundary20211221UpdateRequest) {
	o.Body = body
}

// WithClusterID adds the clusterID to the update params
func (o *UpdateParams) WithClusterID(clusterID string) *UpdateParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the update params
func (o *UpdateParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the update params
func (o *UpdateParams) WithLocationOrganizationID(locationOrganizationID string) *UpdateParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the update params
func (o *UpdateParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the update params
func (o *UpdateParams) WithLocationProjectID(locationProjectID string) *UpdateParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the update params
func (o *UpdateParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
