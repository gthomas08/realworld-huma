package jwtkit

import (
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// GenerateToken creates and signs a JWT with the provided issuer, subject, secret key, and expiration time.
func GenerateToken(issuer string, subject string, secretKey string, ttlSeconds int) (string, error) {
	// Calculate the expiration time
	now := time.Now()
	expiration := now.Add(time.Second * time.Duration(ttlSeconds))

	// Build the JWT
	tok, err := jwt.NewBuilder().
		Issuer(issuer).
		IssuedAt(now).
		Subject(subject).
		Expiration(expiration).
		Build()
	if err != nil {
		return "", err
	}

	// Sign the JWT with the specified algorithm and key
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, []byte(secretKey)))
	if err != nil {
		return "", err
	}

	return string(signed), nil
}

// ValidateToken parses and verifies a JWT with the specified secret key.
func ValidateToken(token string, secretKey string) (jwt.Token, error) {
	return jwt.Parse([]byte(token), jwt.WithKey(jwa.HS256, []byte(secretKey)))
}
