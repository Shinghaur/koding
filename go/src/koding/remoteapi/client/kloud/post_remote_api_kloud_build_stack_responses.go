package kloud

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIKloudBuildStackReader is a Reader for the PostRemoteAPIKloudBuildStack structure.
type PostRemoteAPIKloudBuildStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIKloudBuildStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIKloudBuildStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIKloudBuildStackUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIKloudBuildStackOK creates a PostRemoteAPIKloudBuildStackOK with default headers values
func NewPostRemoteAPIKloudBuildStackOK() *PostRemoteAPIKloudBuildStackOK {
	return &PostRemoteAPIKloudBuildStackOK{}
}

/*PostRemoteAPIKloudBuildStackOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIKloudBuildStackOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIKloudBuildStackOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.buildStack][%d] postRemoteApiKloudBuildStackOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIKloudBuildStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIKloudBuildStackUnauthorized creates a PostRemoteAPIKloudBuildStackUnauthorized with default headers values
func NewPostRemoteAPIKloudBuildStackUnauthorized() *PostRemoteAPIKloudBuildStackUnauthorized {
	return &PostRemoteAPIKloudBuildStackUnauthorized{}
}

/*PostRemoteAPIKloudBuildStackUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIKloudBuildStackUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIKloudBuildStackUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/Kloud.buildStack][%d] postRemoteApiKloudBuildStackUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIKloudBuildStackUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
