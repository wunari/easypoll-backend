// Code generated by go-swagger; DO NOT EDIT.

package poll

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"

	models "github.com/wunari/easypoll-backend/models"
)

// UpdatePollByIDHandlerFunc turns a function with the right signature into a update poll by Id handler
type UpdatePollByIDHandlerFunc func(UpdatePollByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePollByIDHandlerFunc) Handle(params UpdatePollByIDParams) middleware.Responder {
	return fn(params)
}

// UpdatePollByIDHandler interface for that can handle valid update poll by Id params
type UpdatePollByIDHandler interface {
	Handle(UpdatePollByIDParams) middleware.Responder
}

// NewUpdatePollByID creates a new http.Handler for the update poll by Id operation
func NewUpdatePollByID(ctx *middleware.Context, handler UpdatePollByIDHandler) *UpdatePollByID {
	return &UpdatePollByID{Context: ctx, Handler: handler}
}

/*UpdatePollByID swagger:route PUT /polls/{id} poll updatePollById

Update a poll by id

Updates the whole poll object with a new one

Unspecified optional fields will be counted as zero-value and will be overwritten


*/
type UpdatePollByID struct {
	Context *middleware.Context
	Handler UpdatePollByIDHandler
}

func (o *UpdatePollByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdatePollByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdatePollByIDBody update poll by ID body
// swagger:model UpdatePollByIDBody
type UpdatePollByIDBody struct {
	models.PollRequest
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *UpdatePollByIDBody) UnmarshalJSON(raw []byte) error {
	// UpdatePollByIDParamsBodyAO0
	var updatePollByIDParamsBodyAO0 models.PollRequest
	if err := swag.ReadJSON(raw, &updatePollByIDParamsBodyAO0); err != nil {
		return err
	}
	o.PollRequest = updatePollByIDParamsBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o UpdatePollByIDBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	updatePollByIDParamsBodyAO0, err := swag.WriteJSON(o.PollRequest)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, updatePollByIDParamsBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update poll by ID body
func (o *UpdatePollByIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.PollRequest
	if err := o.PollRequest.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *UpdatePollByIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdatePollByIDBody) UnmarshalBinary(b []byte) error {
	var res UpdatePollByIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}