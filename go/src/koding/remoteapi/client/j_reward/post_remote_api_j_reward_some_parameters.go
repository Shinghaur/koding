package j_reward

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

// NewPostRemoteAPIJRewardSomeParams creates a new PostRemoteAPIJRewardSomeParams object
// with the default values initialized.
func NewPostRemoteAPIJRewardSomeParams() *PostRemoteAPIJRewardSomeParams {
	var ()
	return &PostRemoteAPIJRewardSomeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJRewardSomeParamsWithTimeout creates a new PostRemoteAPIJRewardSomeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJRewardSomeParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJRewardSomeParams {
	var ()
	return &PostRemoteAPIJRewardSomeParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJRewardSomeParamsWithContext creates a new PostRemoteAPIJRewardSomeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJRewardSomeParamsWithContext(ctx context.Context) *PostRemoteAPIJRewardSomeParams {
	var ()
	return &PostRemoteAPIJRewardSomeParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJRewardSomeParams contains all the parameters to send to the API endpoint
for the post remote API j reward some operation typically these are written to a http.Request
*/
type PostRemoteAPIJRewardSomeParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJRewardSomeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) WithContext(ctx context.Context) *PostRemoteAPIJRewardSomeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJRewardSomeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j reward some params
func (o *PostRemoteAPIJRewardSomeParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJRewardSomeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
