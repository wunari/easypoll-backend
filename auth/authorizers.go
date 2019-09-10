package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	errors "github.com/go-openapi/errors"

	"github.com/wunari/easypoll-backend/docs/models"
)

// this should be changed to an env variable
var hmacSampleSecret = []byte("OiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9")

// IsValidToken validates if the Bearer token sent by the user is valid
func IsValidToken(token string) (*models.User, error) {
	claims, err := parseAndCheckToken(token)
	if err == nil {
		return &models.User{Name: claims["foo"].(string)}, nil
	}
	return nil, errors.New(401, "the provided token is not valid")
}

func parseAndCheckToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err == nil {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			return claims, nil
		}
	}

	return nil, err
}
