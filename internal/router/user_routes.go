package router

import (
	"github.com/go-chi/chi"
	"github.com/olehhuss/rssagg/internal/handler"
)

// SetupUserRoutes configures user-related routes
func SetupUserRoutes(router chi.Router, userHandler handler.UserHandlerInterface) {
	if userHandler == nil {
		return
	}

	router.Route("/users", func(r chi.Router) {
		r.Get("/me", userHandler.GetUser) // Get current user info
		// Future user routes:
		// r.Put("/me", userHandler.UpdateUser)
		// r.Delete("/me", userHandler.DeleteUser)
	})
}
