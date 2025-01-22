package auth

import (
	"BudgetApp/helpers"
	"BudgetApp/models"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// Define a secret key for signing tokens - in production, this should be in env variables
var jwtSecret = []byte(os.Getenv("APP_SECRET"))

// Claims struct to define JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for the given user
func GenerateToken(user models.User) (string, error) {
	// Set expiration time (e.g., 24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create claims with user data
	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

// ValidateToken verifies and parses the JWT token
func ValidateToken(tokenString string) (*models.User, error) {
	claims := &Claims{}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	db, err := helpers.ConnectToSQLite()

	var user models.User

	db.First(&user, "id = ?", claims.UserID)

	return &user, err
}
