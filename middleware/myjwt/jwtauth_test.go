package myjwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestGenerateJwtToken(t *testing.T) {
	customClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 2, 0)),
			Issuer:    "my-app",
		},
		Role:      "user",
		Platform:  "web",
		SecretKey: []byte("secret-key"),
	}
	tokenString, err := customClaims.GenerateJwtToken()
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
	if tokenString == "" {
		t.Error("Expected token string, but got empty string")
	}
	// Verify the token using the same secret key
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return customClaims.SecretKey, nil
	})
	if err != nil {
		t.Errorf("Unexpected error while verifying token: %v", err)
	}
	if _, ok := token.Claims.(*CustomClaims); !ok {
		t.Error("Token claims is not of expected type CustomClaims")
	}
	// Verify the token contains the expected role and platform information
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		t.Error("Failed to parse token claims as CustomClaims type")
	}
	if claims.Role != customClaims.Role {
		t.Errorf("Expected role %v, but got %v", customClaims.Role, claims.Role)
	}
	if claims.Platform != customClaims.Platform {
		t.Errorf("Expected platform %v, but got %v", customClaims.Platform, claims.Platform)
	}
}
