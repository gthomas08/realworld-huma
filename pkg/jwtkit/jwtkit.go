package jwtkit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type UserClaim struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Bio      *string   `json:"bio"`
	Image    *string   `json:"image"`
}

// GenerateToken creates and signs a JWT with the provided issuer, subject, secret key, and expiration time.
func GenerateToken(issuer string, subject string, secretKey string, ttlSeconds int, claims *UserClaim) (string, error) {
	// Calculate the expiration time
	now := time.Now()
	expiration := now.Add(time.Second * time.Duration(ttlSeconds))

	// Build the JWT
	tok, err := jwt.NewBuilder().
		Issuer(issuer).
		IssuedAt(now).
		Subject(subject).
		Expiration(expiration).
		Claim("user", claims).
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

// ParseUserClaim converts a generic data interface into a UserClaim struct.
// It marshals the input data to JSON format and then unmarshals it into a UserClaim.
func ParseUserClaim(data interface{}) (*UserClaim, error) {
	// Convert the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	// Create an instance of User to unmarshal into
	var user UserClaim

	// Unmarshal the JSON into the struct
	if err := json.Unmarshal(jsonData, &user); err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %w", err)
	}

	return &user, nil
}
