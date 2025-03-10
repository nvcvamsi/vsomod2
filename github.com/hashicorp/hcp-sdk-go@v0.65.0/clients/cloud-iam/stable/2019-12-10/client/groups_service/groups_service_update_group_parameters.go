// Code generated by go-swagger; DO NOT EDIT.

package groups_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
)

// NewGroupsServiceUpdateGroupParams creates a new GroupsServiceUpdateGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGroupsServiceUpdateGroupParams() *GroupsServiceUpdateGroupParams {
	return &GroupsServiceUpdateGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGroupsServiceUpdateGroupParamsWithTimeout creates a new GroupsServiceUpdateGroupParams object
// with the ability to set a timeout on a request.
func NewGroupsServiceUpdateGroupParamsWithTimeout(timeout time.Duration) *GroupsServiceUpdateGroupParams {
	return &GroupsServiceUpdateGroupParams{
		timeout: timeout,
	}
}

// NewGroupsServiceUpdateGroupParamsWithContext creates a new GroupsServiceUpdateGroupParams object
// with the ability to set a context for a request.
func NewGroupsServiceUpdateGroupParamsWithContext(ctx context.Context) *GroupsServiceUpdateGroupParams {
	return &GroupsServiceUpdateGroupParams{
		Context: ctx,
	}
}

// NewGroupsServiceUpdateGroupParamsWithHTTPClient creates a new GroupsServiceUpdateGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewGroupsServiceUpdateGroupParamsWithHTTPClient(client *http.Client) *GroupsServiceUpdateGroupParams {
	return &GroupsServiceUpdateGroupParams{
		HTTPClient: client,
	}
}

/*
GroupsServiceUpdateGroupParams contains all the parameters to send to the API endpoint

	for the groups service update group operation.

	Typically these are written to a http.Request.
*/
type GroupsServiceUpdateGroupParams struct {

	// Body.
	Body *models.HashicorpCloudIamUpdateGroupRequest

	/* ResourceName.

	     resource_name is the resource name of the group.
	iam/organization/<org_id>/group/<group_id>
	*/
	ResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the groups service update group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceUpdateGroupParams) WithDefaults() *GroupsServiceUpdateGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the groups service update group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GroupsServiceUpdateGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) WithTimeout(timeout time.Duration) *GroupsServiceUpdateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) WithContext(ctx context.Context) *GroupsServiceUpdateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) WithHTTPClient(client *http.Client) *GroupsServiceUpdateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) WithBody(body *models.HashicorpCloudIamUpdateGroupRequest) *GroupsServiceUpdateGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) SetBody(body *models.HashicorpCloudIamUpdateGroupRequest) {
	o.Body = body
}

// WithResourceName adds the resourceName to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) WithResourceName(resourceName string) *GroupsServiceUpdateGroupParams {
	o.SetResourceName(resourceName)
	return o
}

// SetResourceName adds the resourceName to the groups service update group params
func (o *GroupsServiceUpdateGroupParams) SetResourceName(resourceName string) {
	o.ResourceName = resourceName
}

// WriteToRequest writes these params to a swagger request
func (o *GroupsServiceUpdateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param resource_name
	if err := r.SetPathParam("resource_name", o.ResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
