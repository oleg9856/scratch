package router

import (
	"github.com/go-chi/chi"
	"github.com/olehhuss/rssagg/internal/handler"
)

// SetupFeedRoutes configures feed-related routes
func SetupFeedRoutes(router chi.Router, feedHandler handler.FeedHandlerInterface) {
	if feedHandler == nil {
		return
	}

	router.Route("/feeds", func(r chi.Router) {
		r.Post("/", feedHandler.CreateFeed)    // Create new feed
		r.Get("/", feedHandler.GetUserFeeds)   // Get user's feeds
		r.Get("/all", feedHandler.GetAllFeeds) // Get all feeds (admin)

		// Feed-specific routes
		r.Route("/{feedID}", func(r chi.Router) {
			// Future feed routes:
			// r.Get("/", feedHandler.GetFeed)          // Get specific feed
			// r.Put("/", feedHandler.UpdateFeed)       // Update feed
			// r.Delete("/", feedHandler.DeleteFeed)    // Delete feed
			// r.Post("/refresh", feedHandler.RefreshFeed) // Manually refresh feed
		})
	})
}
