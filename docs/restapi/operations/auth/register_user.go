// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "github.com/wunari/easypoll-backend/docs/models"
)

// RegisterUserHandlerFunc turns a function with the right signature into a register user handler
type RegisterUserHandlerFunc func(RegisterUserParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RegisterUserHandlerFunc) Handle(params RegisterUserParams) middleware.Responder {
	return fn(params)
}

// RegisterUserHandler interface for that can handle valid register user params
type RegisterUserHandler interface {
	Handle(RegisterUserParams) middleware.Responder
}

// NewRegisterUser creates a new http.Handler for the register user operation
func NewRegisterUser(ctx *middleware.Context, handler RegisterUserHandler) *RegisterUser {
	return &RegisterUser{Context: ctx, Handler: handler}
}

/*RegisterUser swagger:route POST /register auth registerUser

Create a new account

Create a new account and returns a token to be used in requests

*/
type RegisterUser struct {
	Context *middleware.Context
	Handler RegisterUserHandler
}

func (o *RegisterUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewRegisterUserParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// RegisterUserBody register user body
// swagger:model RegisterUserBody
type RegisterUserBody struct {

	// Email of the account
	// Required: true
	Email *string `json:"email"`

	// Name of the user
	// Required: true
	Name *string `json:"name"`

	// Password of the account
	// Required: true
	// Format: password
	Password *strfmt.Password `json:"password"`
}

// Validate validates this register user body
func (o *RegisterUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterUserBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *RegisterUserBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *RegisterUserBody) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"password", "body", o.Password); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"password", "body", "password", o.Password.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterUserBody) UnmarshalBinary(b []byte) error {
	var res RegisterUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// RegisterUserOKBody register user o k body
// swagger:model RegisterUserOKBody
type RegisterUserOKBody struct {
	models.Token

	// user
	User *models.User `json:"user,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *RegisterUserOKBody) UnmarshalJSON(raw []byte) error {
	// RegisterUserOKBodyAO0
	var registerUserOKBodyAO0 models.Token
	if err := swag.ReadJSON(raw, &registerUserOKBodyAO0); err != nil {
		return err
	}
	o.Token = registerUserOKBodyAO0

	// RegisterUserOKBodyAO1
	var dataRegisterUserOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataRegisterUserOKBodyAO1); err != nil {
		return err
	}

	o.User = dataRegisterUserOKBodyAO1.User

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o RegisterUserOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	registerUserOKBodyAO0, err := swag.WriteJSON(o.Token)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, registerUserOKBodyAO0)

	var dataRegisterUserOKBodyAO1 struct {
		User *models.User `json:"user,omitempty"`
	}

	dataRegisterUserOKBodyAO1.User = o.User

	jsonDataRegisterUserOKBodyAO1, errRegisterUserOKBodyAO1 := swag.WriteJSON(dataRegisterUserOKBodyAO1)
	if errRegisterUserOKBodyAO1 != nil {
		return nil, errRegisterUserOKBodyAO1
	}
	_parts = append(_parts, jsonDataRegisterUserOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this register user o k body
func (o *RegisterUserOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Token
	if err := o.Token.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RegisterUserOKBody) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(o.User) { // not required
		return nil
	}

	if o.User != nil {
		if err := o.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registerUserOK" + "." + "user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RegisterUserOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RegisterUserOKBody) UnmarshalBinary(b []byte) error {
	var res RegisterUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
