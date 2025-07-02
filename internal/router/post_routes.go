package router

import (
	"github.com/go-chi/chi"
	"github.com/olehhuss/rssagg/internal/handler"
)

// SetupPostRoutes configures post-related routes
func SetupPostRoutes(router chi.Router, postHandler handler.PostHandlerInterface) {
	if postHandler == nil {
		return
	}

	router.Route("/posts", func(r chi.Router) {
		r.Get("/", postHandler.GetUserPosts) // Get user's posts

		// Future post routes:
		// r.Get("/{postID}", postHandler.GetPost)     // Get specific post
		// r.Post("/{postID}/read", postHandler.MarkAsRead) // Mark post as read
	})
}
