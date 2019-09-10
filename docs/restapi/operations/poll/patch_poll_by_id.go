// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PatchPollByIDHandlerFunc turns a function with the right signature into a patch poll by Id handler
type PatchPollByIDHandlerFunc func(PatchPollByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchPollByIDHandlerFunc) Handle(params PatchPollByIDParams) middleware.Responder {
	return fn(params)
}

// PatchPollByIDHandler interface for that can handle valid patch poll by Id params
type PatchPollByIDHandler interface {
	Handle(PatchPollByIDParams) middleware.Responder
}

// NewPatchPollByID creates a new http.Handler for the patch poll by Id operation
func NewPatchPollByID(ctx *middleware.Context, handler PatchPollByIDHandler) *PatchPollByID {
	return &PatchPollByID{Context: ctx, Handler: handler}
}

/*PatchPollByID swagger:route PATCH /polls/{id} poll patchPollById

Update (part of) a poll by id

Updates (part of) a poll properties, all fields are optional

Unspecified fields will be ignored and won't be updated


*/
type PatchPollByID struct {
	Context *middleware.Context
	Handler PatchPollByIDHandler
}

func (o *PatchPollByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchPollByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
