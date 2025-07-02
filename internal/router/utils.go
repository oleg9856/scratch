package router

import (
	"encoding/json"
	"net/http"
)

// Middleware utilities and common route handlers

// JSONResponse is a helper function for sending JSON responses
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		}
	}
}

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// SendError sends a standardized error response
func SendError(w http.ResponseWriter, status int, message string) {
	response := ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
		Code:    status,
	}
	JSONResponse(w, status, response)
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Service   string `json:"service"`
	Version   string `json:"version,omitempty"`
	Timestamp string `json:"timestamp"`
}

// Future middleware can be added here:
// - Rate limiting middleware
// - Request ID middleware
// - Logging middleware
// - Metrics middleware
