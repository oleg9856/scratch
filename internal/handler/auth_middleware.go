package handler

import (
	"net/http"
	"strings"

	"github.com/olehhuss/rssagg/internal/usecase"
)

// AuthMiddleware validates API key from request headers
type AuthMiddleware struct {
	userService usecase.UserServiceInterface // ← Інтерфейс
}

// NewAuthMiddleware creates new auth middleware
func NewAuthMiddleware(userService usecase.UserServiceInterface) AuthMiddlewareInterface {
	return &AuthMiddleware{userService: userService}
}

// Authenticate validates API key and sets user in context
func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "API key required")
			return
		}

		// Extract API key from "ApiKey <key>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "ApiKey" {
			respondWithError(w, http.StatusUnauthorized, "Invalid authorization format")
			return
		}

		user, err := m.userService.GetUserByAPIKey(r.Context(), parts[1])
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid API key")
			return
		}

		// Add user to request context
		ctx := r.Context()
		ctx = setUserInContext(ctx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
