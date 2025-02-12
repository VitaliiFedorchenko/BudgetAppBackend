package middlewares

import (
	serverUtils "BudgetApp/cmd/utils"
	"BudgetApp/internal/utils"
	"BudgetApp/models"
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

// RoleMiddleware creates a middleware that checks for specific roles
func RoleMiddleware(allowedRoles ...models.UserRole) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// First, ensure authentication
			authMiddleware := AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
				// Get the user from context (which was added by AuthMiddleware)
				user, ok := GetUserFromContext(r)
				if !ok {
					utils.NewResponse(w).ResponseJSON("User not found", http.StatusUnauthorized)
					return
				}

				// Check if user's role is in the allowed roles
				roleAllowed := false
				for _, allowedRole := range allowedRoles {
					if user.Role == allowedRole {
						roleAllowed = true
						break
					}
				}

				if !roleAllowed {
					utils.NewResponse(w).ResponseJSON("Insufficient permissions", http.StatusForbidden)
					return
				}

				// If role is allowed, call the next handler
				next.ServeHTTP(w, r)
			})

			// Run the auth middleware first
			authMiddleware.ServeHTTP(w, r)
		}
	}
}

// GetUserFromContext retrieves the user from the request context
func GetUserFromContext(r *http.Request) (*models.User, bool) {
	user, ok := r.Context().Value(UserContextKey).(*models.User)
	return user, ok
}
