package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccessTokenClaims struct {
	ID          uint      `json:"user_id"`
	Email       string    `json:"email"`
	ClientAppId uuid.UUID `json:"client_app_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	ClientAppId uuid.UUID `json:"client_app_id"`
	jwt.RegisteredClaims
}

func CreateAccessToken(claims AccessTokenClaims) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	// Set standard claims
	expirationTime := time.Minute * 30
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	// Generate client app id
	clientAppId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	claims.ClientAppId = clientAppId

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateRefreshToken(claims RefreshTokenClaims) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	// Set standard claims
	expirationTime := time.Hour * 360
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expirationTime))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	// Generate client app id
	clientAppId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	claims.ClientAppId = clientAppId

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
