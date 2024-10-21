package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Ensure to keep this secure!

// GenerateJWT creates a new JWT token for a given user ID
func GenerateJWT(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   string(userID),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
