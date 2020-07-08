// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserbyIDHandlerFunc turns a function with the right signature into a update userby ID handler
type UpdateUserbyIDHandlerFunc func(UpdateUserbyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserbyIDHandlerFunc) Handle(params UpdateUserbyIDParams) middleware.Responder {
	return fn(params)
}

// UpdateUserbyIDHandler interface for that can handle valid update userby ID params
type UpdateUserbyIDHandler interface {
	Handle(UpdateUserbyIDParams) middleware.Responder
}

// NewUpdateUserbyID creates a new http.Handler for the update userby ID operation
func NewUpdateUserbyID(ctx *middleware.Context, handler UpdateUserbyIDHandler) *UpdateUserbyID {
	return &UpdateUserbyID{Context: ctx, Handler: handler}
}

/*UpdateUserbyID swagger:route PATCH /users/{id} users updateUserbyId

Updates a user by ID.

It updates the user information details by ID


*/
type UpdateUserbyID struct {
	Context *middleware.Context
	Handler UpdateUserbyIDHandler
}

func (o *UpdateUserbyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserbyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}