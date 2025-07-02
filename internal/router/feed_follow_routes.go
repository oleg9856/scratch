package router

import (
	"github.com/go-chi/chi"
	"github.com/olehhuss/rssagg/internal/handler"
)

// SetupFeedFollowRoutes configures feed subscription routes
func SetupFeedFollowRoutes(router chi.Router, followHandler handler.FeedFollowHandlerInterface) {
	if followHandler == nil {
		return
	}

	router.Route("/feeds/{feedID}", func(r chi.Router) {
		r.Post("/follow", followHandler.FollowFeed)     // Subscribe to feed
		r.Delete("/follow", followHandler.UnfollowFeed) // Unsubscribe from feed
	})

	// User's subscriptions
	router.Get("/follows", followHandler.GetUserFollows) // Get user's subscriptions
}
