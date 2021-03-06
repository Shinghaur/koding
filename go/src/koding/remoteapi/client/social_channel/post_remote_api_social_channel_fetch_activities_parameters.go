package social_channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPISocialChannelFetchActivitiesParams creates a new PostRemoteAPISocialChannelFetchActivitiesParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelFetchActivitiesParams() *PostRemoteAPISocialChannelFetchActivitiesParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelFetchActivitiesParamsWithTimeout creates a new PostRemoteAPISocialChannelFetchActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelFetchActivitiesParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchActivitiesParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivitiesParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelFetchActivitiesParamsWithContext creates a new PostRemoteAPISocialChannelFetchActivitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelFetchActivitiesParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchActivitiesParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivitiesParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelFetchActivitiesParams contains all the parameters to send to the API endpoint
for the post remote API social channel fetch activities operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelFetchActivitiesParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchActivitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchActivitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) WithBody(body models.DefaultSelector) *PostRemoteAPISocialChannelFetchActivitiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel fetch activities params
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelFetchActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
