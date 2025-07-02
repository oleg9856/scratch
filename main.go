package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/olehhuss/rssagg/internal/database"
	"github.com/olehhuss/rssagg/internal/handler"
	"github.com/olehhuss/rssagg/internal/infrastructure"
	repository "github.com/olehhuss/rssagg/internal/repository/postgres"
	"github.com/olehhuss/rssagg/internal/usecase"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Load configuration
	config, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := infrastructure.NewDatabase(&config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize SQLC queries
	queries := database.New(db)

	// Initialize repositories
	userRepo := repository.NewUserRepository(queries)
	// feedRepo := repository.NewFeedRepository(queries) // TODO: implement later

	// Initialize use cases
	userService := usecase.NewUserService(userRepo)
	// feedService := usecase.NewFeedService(feedRepo, userRepo) // TODO: implement later

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	authMiddleware := handler.NewAuthMiddleware(userService)

	// Setup router
	router := setupRouter(userHandler, authMiddleware)

	// Start server
	log.Printf("Server starting on port %s", config.Server.Port)
	if err := http.ListenAndServe(":"+config.Server.Port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func setupRouter(userHandler handler.UserHandlerInterface, authMiddleware handler.AuthMiddlewareInterface) chi.Router {
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

	// Public routes
	v1Router.Get("/healthz", handler.HealthCheck)
	v1Router.Get("/error", handler.ErrorTest)
	v1Router.Post("/users", userHandler.CreateUser)

	// Protected routes (require authentication)
	v1Router.Group(func(r chi.Router) {
		// Add auth middleware to protected routes
		// r.Use(authMiddleware.Authenticate)

		// Future protected endpoints will go here:
		// r.Post("/feeds", feedHandler.CreateFeed)
		// r.Get("/feeds", feedHandler.GetUserFeeds)
		// r.Get("/posts", postHandler.GetUserPosts)
	})

	router.Mount("/v1", v1Router)
	return router
}
