package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	errors "github.com/go-openapi/errors"

	"github.com/wunari/easypoll-backend/docs/models"
)

// IsValidToken validates if the Bearer token sent by the user is valid
func IsValidToken(token string) (*models.User, error) {
	claims, err := parseAndCheckToken(token)
	if err == nil {
		return &models.User{Name: claims["name"].(string)}, nil
	}
	return nil, errors.New(401, "the provided token is not valid")
}

func parseAndCheckToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", parsedToken.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err == nil {
		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			return claims, nil
		}
	}

	return nil, err
}
