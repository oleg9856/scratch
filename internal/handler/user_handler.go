package handler

import (
	"encoding/json"
	"net/http"

	"github.com/olehhuss/rssagg/internal/interfaces"
)

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userService interfaces.UserServiceInterface
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService interfaces.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser handles POST /user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req interfaces.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	user, err := h.userService.CreateUser(r.Context(), req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// GetUser handles GET /users/me - get current user info
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, ok := getUserFromContext(r.Context())
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "User not found in context")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
