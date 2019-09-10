package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/auth"
)

// GetAuthenticatedUserHandlerFunc return the authenticated user account
func GetAuthenticatedUserHandlerFunc(params auth.GetAuthenticatedUserParams, user interface{}) middleware.Responder {
	return auth.NewGetAuthenticatedUserOK()
}
