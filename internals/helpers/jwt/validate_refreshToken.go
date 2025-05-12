package jwt

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateRefreshToken(tokenString string) (bool, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
