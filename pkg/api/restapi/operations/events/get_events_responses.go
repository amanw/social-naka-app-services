// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
)

// GetEventsOKCode is the HTTP code returned for type GetEventsOK
const GetEventsOKCode int = 200

/*GetEventsOK OK

swagger:response getEventsOK
*/
type GetEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.EventResponse `json:"body,omitempty"`
}

// NewGetEventsOK creates GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {

	return &GetEventsOK{}
}

// WithPayload adds the payload to the get events o k response
func (o *GetEventsOK) WithPayload(payload *models.EventResponse) *GetEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events o k response
func (o *GetEventsOK) SetPayload(payload *models.EventResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsBadRequestCode is the HTTP code returned for type GetEventsBadRequest
const GetEventsBadRequestCode int = 400

/*GetEventsBadRequest The api is Unauthorized

swagger:response getEventsBadRequest
*/
type GetEventsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetEventsBadRequest creates GetEventsBadRequest with default headers values
func NewGetEventsBadRequest() *GetEventsBadRequest {

	return &GetEventsBadRequest{}
}

// WithPayload adds the payload to the get events bad request response
func (o *GetEventsBadRequest) WithPayload(payload *models.Response) *GetEventsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events bad request response
func (o *GetEventsBadRequest) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsUnauthorizedCode is the HTTP code returned for type GetEventsUnauthorized
const GetEventsUnauthorizedCode int = 401

/*GetEventsUnauthorized The api is Unauthorized

swagger:response getEventsUnauthorized
*/
type GetEventsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetEventsUnauthorized creates GetEventsUnauthorized with default headers values
func NewGetEventsUnauthorized() *GetEventsUnauthorized {

	return &GetEventsUnauthorized{}
}

// WithPayload adds the payload to the get events unauthorized response
func (o *GetEventsUnauthorized) WithPayload(payload *models.Response) *GetEventsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events unauthorized response
func (o *GetEventsUnauthorized) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsNotFoundCode is the HTTP code returned for type GetEventsNotFound
const GetEventsNotFoundCode int = 404

/*GetEventsNotFound The api is not found

swagger:response getEventsNotFound
*/
type GetEventsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetEventsNotFound creates GetEventsNotFound with default headers values
func NewGetEventsNotFound() *GetEventsNotFound {

	return &GetEventsNotFound{}
}

// WithPayload adds the payload to the get events not found response
func (o *GetEventsNotFound) WithPayload(payload *models.Response) *GetEventsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events not found response
func (o *GetEventsNotFound) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsInternalServerErrorCode is the HTTP code returned for type GetEventsInternalServerError
const GetEventsInternalServerErrorCode int = 500

/*GetEventsInternalServerError Internal Server Error

swagger:response getEventsInternalServerError
*/
type GetEventsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewGetEventsInternalServerError creates GetEventsInternalServerError with default headers values
func NewGetEventsInternalServerError() *GetEventsInternalServerError {

	return &GetEventsInternalServerError{}
}

// WithPayload adds the payload to the get events internal server error response
func (o *GetEventsInternalServerError) WithPayload(payload *models.Response) *GetEventsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events internal server error response
func (o *GetEventsInternalServerError) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
