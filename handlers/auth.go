package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/auth"
)

// this should be changed to an env variable
var hmacSampleSecret = []byte("OiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9")

// in-memory "database", it will be changed to proper database later on
var users = []*models.User{}

// LoginUserHandlerFunc authenticates an user
func LoginUserHandlerFunc(params auth.LoginUserParams) middleware.Responder {
	// find user
	for _, user := range users {
		if user.Email == *params.Body.Email {
			if user.Password == *params.Body.Password {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"name":  user.Name,
					"email": user.Email,
					"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				})
				tokenString, _ := token.SignedString(hmacSampleSecret)

				tokenModel := models.Token{Token: tokenString}
				return auth.NewLoginUserOK().WithPayload(&auth.LoginUserOKBody{User: user, Token: tokenModel})
			}
			return auth.NewLoginUserBadRequest().WithPayload(&models.Error{Code: 400, Message: "Invalid password"})
		}
	}

	return auth.NewLoginUserNotFound().WithPayload(&models.Error{Code: 404, Message: "Email not found"})
}

// RegisterUserHandlerFunc creates a new user in the database
func RegisterUserHandlerFunc(params auth.RegisterUserParams) middleware.Responder {
	// store user
	user := models.User{Name: *params.Body.Name, Email: *params.Body.Email, Password: *params.Body.Password}
	users = append(users, &user)

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  params.Body.Name,
		"email": params.Body.Email,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, _ := token.SignedString(hmacSampleSecret)

	tokenModel := models.Token{Token: tokenString}
	return auth.NewRegisterUserOK().WithPayload(&auth.RegisterUserOKBody{User: &user, Token: tokenModel})
}

// GetAuthenticatedUserHandlerFunc return the authenticated user account
func GetAuthenticatedUserHandlerFunc(params auth.GetAuthenticatedUserParams, user *models.User) middleware.Responder {
	return auth.NewGetAuthenticatedUserOK().WithPayload(&auth.GetAuthenticatedUserOKBody{User: *user})
}
