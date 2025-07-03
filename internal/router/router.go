package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/olehhuss/rssagg/internal/handler"
)

// Config holds all handlers and middleware needed for routing
type Config struct {
	HomeHandler       handler.HomeHandlerInterface
	UserHandler       handler.UserHandlerInterface
	FeedHandler       handler.FeedHandlerInterface
	PostHandler       handler.PostHandlerInterface
	FeedFollowHandler handler.FeedFollowHandlerInterface
	AuthMiddleware    handler.AuthMiddlewareInterface
}

// Setup creates and configures the main router
func Setup(config *Config) chi.Router {
	router := chi.NewRouter()

	// CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// API v1 routes
	v1Router := chi.NewRouter()

	// Setup public routes
	setupPublicRoutes(v1Router, config)

	// Setup protected routes
	setupProtectedRoutes(v1Router, config)

	// Setup home page routes (public, but at root level)
	setupHomeRoutes(router, config)

	router.Mount("/v1", v1Router)
	return router
}

// setupPublicRoutes configures routes that don't require authentication
func setupPublicRoutes(router chi.Router, config *Config) {
	// Health check and error testing
	router.Get("/healthz", handler.HealthCheck)
	router.Get("/error", handler.ErrorTest)

	// User registration (public)
	if config.UserHandler != nil {
		router.Post("/users", config.UserHandler.CreateUser)
	}
}

// setupProtectedRoutes configures routes that require authentication
func setupProtectedRoutes(router chi.Router, config *Config) {
	router.Group(func(r chi.Router) {
		// Apply authentication middleware to all routes in this group
		if config.AuthMiddleware != nil {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					config.AuthMiddleware.Authenticate(next.ServeHTTP)(w, req)
				})
			})
		}

		// Setup route groups
		SetupUserRoutes(r, config.UserHandler)
		SetupFeedRoutes(r, config.FeedHandler)
		SetupPostRoutes(r, config.PostHandler)
		SetupFeedFollowRoutes(r, config.FeedFollowHandler)
	})
}

// setupHomeRoutes configures the home/landing page routes (public)
func setupHomeRoutes(router chi.Router, config *Config) {
	if config.HomeHandler != nil {
		// Landing page and info endpoints
		router.Get("/", config.HomeHandler.GetHome)
		router.Get("/about", config.HomeHandler.GetAbout)
		router.Get("/api", config.HomeHandler.GetAPIInfo)
	}
}
