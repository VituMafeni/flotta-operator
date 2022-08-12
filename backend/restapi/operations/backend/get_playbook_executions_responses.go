// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/project-flotta/flotta-operator/backend/models"
	commonmodel "github.com/project-flotta/flotta-operator/models"
)

// GetPlaybookExecutionsOKCode is the HTTP code returned for type GetPlaybookExecutionsOK
const GetPlaybookExecutionsOKCode int = 200

/*GetPlaybookExecutionsOK OK

swagger:response getPlaybookExecutionsOK
*/
type GetPlaybookExecutionsOK struct {

	/*
	  In: Body
	*/
	Payload commonmodel.PlaybookExecutionsResponse `json:"body,omitempty"`
}

// NewGetPlaybookExecutionsOK creates GetPlaybookExecutionsOK with default headers values
func NewGetPlaybookExecutionsOK() *GetPlaybookExecutionsOK {

	return &GetPlaybookExecutionsOK{}
}

// WithPayload adds the payload to the get playbook executions o k response
func (o *GetPlaybookExecutionsOK) WithPayload(payload commonmodel.PlaybookExecutionsResponse) *GetPlaybookExecutionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get playbook executions o k response
func (o *GetPlaybookExecutionsOK) SetPayload(payload commonmodel.PlaybookExecutionsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlaybookExecutionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetPlaybookExecutionsUnauthorizedCode is the HTTP code returned for type GetPlaybookExecutionsUnauthorized
const GetPlaybookExecutionsUnauthorizedCode int = 401

/*GetPlaybookExecutionsUnauthorized Unauthorized

swagger:response getPlaybookExecutionsUnauthorized
*/
type GetPlaybookExecutionsUnauthorized struct {
}

// NewGetPlaybookExecutionsUnauthorized creates GetPlaybookExecutionsUnauthorized with default headers values
func NewGetPlaybookExecutionsUnauthorized() *GetPlaybookExecutionsUnauthorized {

	return &GetPlaybookExecutionsUnauthorized{}
}

// WriteResponse to the client
func (o *GetPlaybookExecutionsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetPlaybookExecutionsForbiddenCode is the HTTP code returned for type GetPlaybookExecutionsForbidden
const GetPlaybookExecutionsForbiddenCode int = 403

/*GetPlaybookExecutionsForbidden Forbidden

swagger:response getPlaybookExecutionsForbidden
*/
type GetPlaybookExecutionsForbidden struct {
}

// NewGetPlaybookExecutionsForbidden creates GetPlaybookExecutionsForbidden with default headers values
func NewGetPlaybookExecutionsForbidden() *GetPlaybookExecutionsForbidden {

	return &GetPlaybookExecutionsForbidden{}
}

// WriteResponse to the client
func (o *GetPlaybookExecutionsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

/*GetPlaybookExecutionsDefault Error

swagger:response getPlaybookExecutionsDefault
*/
type GetPlaybookExecutionsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetPlaybookExecutionsDefault creates GetPlaybookExecutionsDefault with default headers values
func NewGetPlaybookExecutionsDefault(code int) *GetPlaybookExecutionsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetPlaybookExecutionsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get playbook executions default response
func (o *GetPlaybookExecutionsDefault) WithStatusCode(code int) *GetPlaybookExecutionsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get playbook executions default response
func (o *GetPlaybookExecutionsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get playbook executions default response
func (o *GetPlaybookExecutionsDefault) WithPayload(payload *models.Error) *GetPlaybookExecutionsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get playbook executions default response
func (o *GetPlaybookExecutionsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPlaybookExecutionsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}