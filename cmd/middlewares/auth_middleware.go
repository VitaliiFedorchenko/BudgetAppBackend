package middlewares

import (
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/utils"
	"context"
	"net/http"
	"strings"
)

// ContextKey is a custom type for context keys to avoid collisions
type ContextKey string

// UserContextKey is the key used to store the user in the context
const UserContextKey ContextKey = "user"

// AuthMiddleware verifies the JWT token and adds the user to the request context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")

		// Check if Authorization header exists
		if authHeader == "" {
			utils.NewResponse(w).ResponseJSON("No authorization header provided", http.StatusUnauthorized)
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.NewResponse(w).ResponseJSON("Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Get user from token and validate the token
		user, err := serverUtils.GetUserFromAuthToken(r)
		if err != nil {
			utils.NewResponse(w).ResponseJSON("Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Create a new context with the user using the custom key type
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserContextKey, user)

		// Call the next handler with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
