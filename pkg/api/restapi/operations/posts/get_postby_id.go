// Code generated by go-swagger; DO NOT EDIT.

package posts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetPostbyIDHandlerFunc turns a function with the right signature into a get postby ID handler
type GetPostbyIDHandlerFunc func(GetPostbyIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPostbyIDHandlerFunc) Handle(params GetPostbyIDParams) middleware.Responder {
	return fn(params)
}

// GetPostbyIDHandler interface for that can handle valid get postby ID params
type GetPostbyIDHandler interface {
	Handle(GetPostbyIDParams) middleware.Responder
}

// NewGetPostbyID creates a new http.Handler for the get postby ID operation
func NewGetPostbyID(ctx *middleware.Context, handler GetPostbyIDHandler) *GetPostbyID {
	return &GetPostbyID{Context: ctx, Handler: handler}
}

/*GetPostbyID swagger:route GET /posts/{id} posts getPostbyId

Gets a post by ID.

It gets the post information details by ID


*/
type GetPostbyID struct {
	Context *middleware.Context
	Handler GetPostbyIDHandler
}

func (o *GetPostbyID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPostbyIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
