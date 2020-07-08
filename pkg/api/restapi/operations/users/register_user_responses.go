// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/amanw/social-naka-app-services/pkg/api/models"
)

// RegisterUserOKCode is the HTTP code returned for type RegisterUserOK
const RegisterUserOKCode int = 200

/*RegisterUserOK OK

swagger:response registerUserOK
*/
type RegisterUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewRegisterUserOK creates RegisterUserOK with default headers values
func NewRegisterUserOK() *RegisterUserOK {

	return &RegisterUserOK{}
}

// WithPayload adds the payload to the register user o k response
func (o *RegisterUserOK) WithPayload(payload *models.User) *RegisterUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user o k response
func (o *RegisterUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserBadRequestCode is the HTTP code returned for type RegisterUserBadRequest
const RegisterUserBadRequestCode int = 400

/*RegisterUserBadRequest The api is Unauthorized

swagger:response registerUserBadRequest
*/
type RegisterUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserBadRequest creates RegisterUserBadRequest with default headers values
func NewRegisterUserBadRequest() *RegisterUserBadRequest {

	return &RegisterUserBadRequest{}
}

// WithPayload adds the payload to the register user bad request response
func (o *RegisterUserBadRequest) WithPayload(payload *models.Error) *RegisterUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user bad request response
func (o *RegisterUserBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserUnauthorizedCode is the HTTP code returned for type RegisterUserUnauthorized
const RegisterUserUnauthorizedCode int = 401

/*RegisterUserUnauthorized The api is Unauthorized

swagger:response registerUserUnauthorized
*/
type RegisterUserUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserUnauthorized creates RegisterUserUnauthorized with default headers values
func NewRegisterUserUnauthorized() *RegisterUserUnauthorized {

	return &RegisterUserUnauthorized{}
}

// WithPayload adds the payload to the register user unauthorized response
func (o *RegisterUserUnauthorized) WithPayload(payload *models.Error) *RegisterUserUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user unauthorized response
func (o *RegisterUserUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserNotFoundCode is the HTTP code returned for type RegisterUserNotFound
const RegisterUserNotFoundCode int = 404

/*RegisterUserNotFound The api is not found

swagger:response registerUserNotFound
*/
type RegisterUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserNotFound creates RegisterUserNotFound with default headers values
func NewRegisterUserNotFound() *RegisterUserNotFound {

	return &RegisterUserNotFound{}
}

// WithPayload adds the payload to the register user not found response
func (o *RegisterUserNotFound) WithPayload(payload *models.Error) *RegisterUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user not found response
func (o *RegisterUserNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserConflictCode is the HTTP code returned for type RegisterUserConflict
const RegisterUserConflictCode int = 409

/*RegisterUserConflict The user already exists

swagger:response registerUserConflict
*/
type RegisterUserConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserConflict creates RegisterUserConflict with default headers values
func NewRegisterUserConflict() *RegisterUserConflict {

	return &RegisterUserConflict{}
}

// WithPayload adds the payload to the register user conflict response
func (o *RegisterUserConflict) WithPayload(payload *models.Error) *RegisterUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user conflict response
func (o *RegisterUserConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RegisterUserInternalServerErrorCode is the HTTP code returned for type RegisterUserInternalServerError
const RegisterUserInternalServerErrorCode int = 500

/*RegisterUserInternalServerError Internal Server Error

swagger:response registerUserInternalServerError
*/
type RegisterUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRegisterUserInternalServerError creates RegisterUserInternalServerError with default headers values
func NewRegisterUserInternalServerError() *RegisterUserInternalServerError {

	return &RegisterUserInternalServerError{}
}

// WithPayload adds the payload to the register user internal server error response
func (o *RegisterUserInternalServerError) WithPayload(payload *models.Error) *RegisterUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register user internal server error response
func (o *RegisterUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
