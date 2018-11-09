// Code generated by go-swagger; DO NOT EDIT.

package predict

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TextsHandlerFunc turns a function with the right signature into a texts handler
type TextsHandlerFunc func(TextsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TextsHandlerFunc) Handle(params TextsParams) middleware.Responder {
	return fn(params)
}

// TextsHandler interface for that can handle valid texts params
type TextsHandler interface {
	Handle(TextsParams) middleware.Responder
}

// NewTexts creates a new http.Handler for the texts operation
func NewTexts(ctx *middleware.Context, handler TextsHandler) *Texts {
	return &Texts{Context: ctx, Handler: handler}
}

/*Texts swagger:route POST /predict/text Predict texts

Text method receives a list base64 encoded texts and runs
the predictor on all the texts.

The result is a prediction feature list for each text.

*/
type Texts struct {
	Context *middleware.Context
	Handler TextsHandler
}

func (o *Texts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTextsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
