// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ModelManifestsHandlerFunc turns a function with the right signature into a model manifests handler
type ModelManifestsHandlerFunc func(ModelManifestsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ModelManifestsHandlerFunc) Handle(params ModelManifestsParams) middleware.Responder {
	return fn(params)
}

// ModelManifestsHandler interface for that can handle valid model manifests params
type ModelManifestsHandler interface {
	Handle(ModelManifestsParams) middleware.Responder
}

// NewModelManifests creates a new http.Handler for the model manifests operation
func NewModelManifests(ctx *middleware.Context, handler ModelManifestsHandler) *ModelManifests {
	return &ModelManifests{Context: ctx, Handler: handler}
}

/*ModelManifests swagger:route GET /registry/models/manifest Registry modelManifests

ModelManifests model manifests API

*/
type ModelManifests struct {
	Context *middleware.Context
	Handler ModelManifestsHandler
}

func (o *ModelManifests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewModelManifestsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
