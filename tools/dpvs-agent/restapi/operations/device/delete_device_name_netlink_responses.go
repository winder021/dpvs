// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteDeviceNameNetlinkOKCode is the HTTP code returned for type DeleteDeviceNameNetlinkOK
const DeleteDeviceNameNetlinkOKCode int = 200

/*
DeleteDeviceNameNetlinkOK Success

swagger:response deleteDeviceNameNetlinkOK
*/
type DeleteDeviceNameNetlinkOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteDeviceNameNetlinkOK creates DeleteDeviceNameNetlinkOK with default headers values
func NewDeleteDeviceNameNetlinkOK() *DeleteDeviceNameNetlinkOK {

	return &DeleteDeviceNameNetlinkOK{}
}

// WithPayload adds the payload to the delete device name netlink o k response
func (o *DeleteDeviceNameNetlinkOK) WithPayload(payload string) *DeleteDeviceNameNetlinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete device name netlink o k response
func (o *DeleteDeviceNameNetlinkOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeviceNameNetlinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteDeviceNameNetlinkInternalServerErrorCode is the HTTP code returned for type DeleteDeviceNameNetlinkInternalServerError
const DeleteDeviceNameNetlinkInternalServerErrorCode int = 500

/*
DeleteDeviceNameNetlinkInternalServerError Not Found

swagger:response deleteDeviceNameNetlinkInternalServerError
*/
type DeleteDeviceNameNetlinkInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteDeviceNameNetlinkInternalServerError creates DeleteDeviceNameNetlinkInternalServerError with default headers values
func NewDeleteDeviceNameNetlinkInternalServerError() *DeleteDeviceNameNetlinkInternalServerError {

	return &DeleteDeviceNameNetlinkInternalServerError{}
}

// WithPayload adds the payload to the delete device name netlink internal server error response
func (o *DeleteDeviceNameNetlinkInternalServerError) WithPayload(payload string) *DeleteDeviceNameNetlinkInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete device name netlink internal server error response
func (o *DeleteDeviceNameNetlinkInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteDeviceNameNetlinkInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}