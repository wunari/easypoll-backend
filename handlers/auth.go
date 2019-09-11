package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/auth"
)

// GetAuthenticatedUserHandlerFunc return the authenticated user account
func GetAuthenticatedUserHandlerFunc(params auth.GetAuthenticatedUserParams, user *models.User) middleware.Responder {
	return auth.NewGetAuthenticatedUserOK().WithPayload(&auth.GetAuthenticatedUserOKBody{User: *user})
}
