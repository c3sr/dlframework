// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// OpenHandlerFunc turns a function with the right signature into a open handler
type OpenHandlerFunc func(OpenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn OpenHandlerFunc) Handle(params OpenParams) middleware.Responder {
	return fn(params)
}

// OpenHandler interface for that can handle valid open params
type OpenHandler interface {
	Handle(OpenParams) middleware.Responder
}

// NewOpen creates a new http.Handler for the open operation
func NewOpen(ctx *middleware.Context, handler OpenHandler) *Open {
	return &Open{Context: ctx, Handler: handler}
}

/*Open swagger:route POST /v1/predict/open Predict open

Opens a predictor and returns an id where the predictor
is accessible. The id can be used to perform inference
requests.

*/
type Open struct {
	Context *middleware.Context
	Handler OpenHandler
}

func (o *Open) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewOpenParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}