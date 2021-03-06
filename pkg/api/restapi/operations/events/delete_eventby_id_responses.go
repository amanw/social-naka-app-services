// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
)

// DeleteEventbyIDOKCode is the HTTP code returned for type DeleteEventbyIDOK
const DeleteEventbyIDOKCode int = 200

/*DeleteEventbyIDOK API Custom Status

swagger:response deleteEventbyIdOK
*/
type DeleteEventbyIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewDeleteEventbyIDOK creates DeleteEventbyIDOK with default headers values
func NewDeleteEventbyIDOK() *DeleteEventbyIDOK {

	return &DeleteEventbyIDOK{}
}

// WithPayload adds the payload to the delete eventby Id o k response
func (o *DeleteEventbyIDOK) WithPayload(payload *models.Response) *DeleteEventbyIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete eventby Id o k response
func (o *DeleteEventbyIDOK) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventbyIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventbyIDBadRequestCode is the HTTP code returned for type DeleteEventbyIDBadRequest
const DeleteEventbyIDBadRequestCode int = 400

/*DeleteEventbyIDBadRequest The api is Unauthorized

swagger:response deleteEventbyIdBadRequest
*/
type DeleteEventbyIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewDeleteEventbyIDBadRequest creates DeleteEventbyIDBadRequest with default headers values
func NewDeleteEventbyIDBadRequest() *DeleteEventbyIDBadRequest {

	return &DeleteEventbyIDBadRequest{}
}

// WithPayload adds the payload to the delete eventby Id bad request response
func (o *DeleteEventbyIDBadRequest) WithPayload(payload *models.Response) *DeleteEventbyIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete eventby Id bad request response
func (o *DeleteEventbyIDBadRequest) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventbyIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventbyIDUnauthorizedCode is the HTTP code returned for type DeleteEventbyIDUnauthorized
const DeleteEventbyIDUnauthorizedCode int = 401

/*DeleteEventbyIDUnauthorized The api is Unauthorized

swagger:response deleteEventbyIdUnauthorized
*/
type DeleteEventbyIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewDeleteEventbyIDUnauthorized creates DeleteEventbyIDUnauthorized with default headers values
func NewDeleteEventbyIDUnauthorized() *DeleteEventbyIDUnauthorized {

	return &DeleteEventbyIDUnauthorized{}
}

// WithPayload adds the payload to the delete eventby Id unauthorized response
func (o *DeleteEventbyIDUnauthorized) WithPayload(payload *models.Response) *DeleteEventbyIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete eventby Id unauthorized response
func (o *DeleteEventbyIDUnauthorized) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventbyIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventbyIDNotFoundCode is the HTTP code returned for type DeleteEventbyIDNotFound
const DeleteEventbyIDNotFoundCode int = 404

/*DeleteEventbyIDNotFound The api is not found

swagger:response deleteEventbyIdNotFound
*/
type DeleteEventbyIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewDeleteEventbyIDNotFound creates DeleteEventbyIDNotFound with default headers values
func NewDeleteEventbyIDNotFound() *DeleteEventbyIDNotFound {

	return &DeleteEventbyIDNotFound{}
}

// WithPayload adds the payload to the delete eventby Id not found response
func (o *DeleteEventbyIDNotFound) WithPayload(payload *models.Response) *DeleteEventbyIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete eventby Id not found response
func (o *DeleteEventbyIDNotFound) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventbyIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteEventbyIDInternalServerErrorCode is the HTTP code returned for type DeleteEventbyIDInternalServerError
const DeleteEventbyIDInternalServerErrorCode int = 500

/*DeleteEventbyIDInternalServerError Internal Server Error

swagger:response deleteEventbyIdInternalServerError
*/
type DeleteEventbyIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Response `json:"body,omitempty"`
}

// NewDeleteEventbyIDInternalServerError creates DeleteEventbyIDInternalServerError with default headers values
func NewDeleteEventbyIDInternalServerError() *DeleteEventbyIDInternalServerError {

	return &DeleteEventbyIDInternalServerError{}
}

// WithPayload adds the payload to the delete eventby Id internal server error response
func (o *DeleteEventbyIDInternalServerError) WithPayload(payload *models.Response) *DeleteEventbyIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete eventby Id internal server error response
func (o *DeleteEventbyIDInternalServerError) SetPayload(payload *models.Response) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteEventbyIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
