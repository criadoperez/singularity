// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/data-preservation-programs/singularity/client/swagger/models"
)

// NewCreateWebdavStorageParams creates a new CreateWebdavStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWebdavStorageParams() *CreateWebdavStorageParams {
	return &CreateWebdavStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWebdavStorageParamsWithTimeout creates a new CreateWebdavStorageParams object
// with the ability to set a timeout on a request.
func NewCreateWebdavStorageParamsWithTimeout(timeout time.Duration) *CreateWebdavStorageParams {
	return &CreateWebdavStorageParams{
		timeout: timeout,
	}
}

// NewCreateWebdavStorageParamsWithContext creates a new CreateWebdavStorageParams object
// with the ability to set a context for a request.
func NewCreateWebdavStorageParamsWithContext(ctx context.Context) *CreateWebdavStorageParams {
	return &CreateWebdavStorageParams{
		Context: ctx,
	}
}

// NewCreateWebdavStorageParamsWithHTTPClient creates a new CreateWebdavStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWebdavStorageParamsWithHTTPClient(client *http.Client) *CreateWebdavStorageParams {
	return &CreateWebdavStorageParams{
		HTTPClient: client,
	}
}

/*
CreateWebdavStorageParams contains all the parameters to send to the API endpoint

	for the create webdav storage operation.

	Typically these are written to a http.Request.
*/
type CreateWebdavStorageParams struct {

	/* Request.

	   Request body
	*/
	Request *models.StorageCreateWebdavStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create webdav storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebdavStorageParams) WithDefaults() *CreateWebdavStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create webdav storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWebdavStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create webdav storage params
func (o *CreateWebdavStorageParams) WithTimeout(timeout time.Duration) *CreateWebdavStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create webdav storage params
func (o *CreateWebdavStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create webdav storage params
func (o *CreateWebdavStorageParams) WithContext(ctx context.Context) *CreateWebdavStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create webdav storage params
func (o *CreateWebdavStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create webdav storage params
func (o *CreateWebdavStorageParams) WithHTTPClient(client *http.Client) *CreateWebdavStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create webdav storage params
func (o *CreateWebdavStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequest adds the request to the create webdav storage params
func (o *CreateWebdavStorageParams) WithRequest(request *models.StorageCreateWebdavStorageRequest) *CreateWebdavStorageParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create webdav storage params
func (o *CreateWebdavStorageParams) SetRequest(request *models.StorageCreateWebdavStorageRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWebdavStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}