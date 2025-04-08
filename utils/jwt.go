package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetUserIDFromToken(tokenString string) (string, error) {
	// Just parse the token, skip signature validation (for now)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
	}

	return "", fmt.Errorf("invalid token or missing sub claim")
}
