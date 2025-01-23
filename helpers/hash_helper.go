package helpers

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes a plain-text password and returns its hashed version
// The cost parameter determines how computationally intensive the hashing will be
func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password cannot be empty")
	}

	// Generate hash with default cost (10)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CheckPassword compares a plain-text password with a hashed password
// Returns nil if they match, error otherwise
func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
