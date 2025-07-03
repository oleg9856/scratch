package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/olehhuss/rssagg/internal/database"
	"github.com/olehhuss/rssagg/internal/handler"
	"github.com/olehhuss/rssagg/internal/infrastructure"
	repository "github.com/olehhuss/rssagg/internal/repository/postgres"
	"github.com/olehhuss/rssagg/internal/router"
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
	feedRepo := repository.NewFeedRepository(queries)

	// Initialize use cases
	userService := usecase.NewUserService(userRepo)
	feedService := usecase.NewFeedService(feedRepo, userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	feedHandler := handler.NewFeedHandler(feedService)
	authMiddleware := handler.NewAuthMiddleware(userService)
	homeHandler := handler.NewHomeHandler(userService, feedService)

	// Create router configuration
	routerConfig := &router.Config{
		HomeHandler:    homeHandler,
		UserHandler:    userHandler,
		FeedHandler:    feedHandler,
		AuthMiddleware: authMiddleware,
		// PostHandler and FeedFollowHandler will be added when implemented
	}

	// Setup router
	r := router.Setup(routerConfig)

	// Start server
	log.Printf("Server starting on port %s", config.Server.Port)
	if err := http.ListenAndServe(":"+config.Server.Port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
