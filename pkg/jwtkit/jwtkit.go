package jwtkit

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(iss string, sub string, secretKey string, seconds int) (string, error) {
	now := time.Now()
	exp := now.Add(time.Second * time.Duration(seconds))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": exp.Unix(),
		"iat": now.Unix(),
		"iss": iss,
		"sub": sub,
	})

	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}
