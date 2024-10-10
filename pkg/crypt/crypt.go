package crypt

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a password string and returns a hashed version of it using bcrypt's
// GenerateFromPassword function with a cost of 14. The hashed password is returned as a
// string, and any error encountered is returned as the second return value.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if a given password matches a given hash. The function
// returns true if the password matches the hash, and false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
