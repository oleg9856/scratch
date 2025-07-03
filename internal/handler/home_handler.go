package handler

import (
	"net/http"
	"time"

	"github.com/olehhuss/rssagg/internal/usecase"
)

// HomeHandler handles requests for the home/landing page
type HomeHandler struct {
	userService usecase.UserServiceInterface
	feedService usecase.FeedServiceInterface
}

// NewHomeHandler creates a new home handler
func NewHomeHandler(userService usecase.UserServiceInterface, feedService usecase.FeedServiceInterface) *HomeHandler {
	return &HomeHandler{
		userService: userService,
		feedService: feedService,
	}
}

// HomeResponse represents the home page data
type HomeResponse struct {
	AppName     string    `json:"app_name"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	Stats       AppStats  `json:"stats"`
	Features    []Feature `json:"features"`
}

type AppStats struct {
	TotalFeeds int `json:"total_feeds"`
	TotalUsers int `json:"total_users"`
	// Додайте більше статистики за потреби
}

type Feature struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon,omitempty"`
}

// GetHome handles GET / - main landing page
func (h *HomeHandler) GetHome(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get real statistics
	totalUsers, err := h.userService.GetUserCount(ctx)
	if err != nil {
		// Log error but continue with 0 count
		totalUsers = 0
	}

	totalFeeds, err := h.feedService.GetFeedCount(ctx)
	if err != nil {
		// Log error but continue with 0 count
		totalFeeds = 0
	}

	stats := AppStats{
		TotalFeeds: totalFeeds,
		TotalUsers: totalUsers,
	}

	features := []Feature{
		{
			Name:        "RSS Feed Aggregation",
			Description: "Collect and organize RSS feeds from multiple sources",
			Icon:        "📰",
		},
		{
			Name:        "Personal Dashboard",
			Description: "View all your feeds in one convenient location",
			Icon:        "📊",
		},
		{
			Name:        "API Access",
			Description: "Full REST API for developers and integrations",
			Icon:        "🔌",
		},
		{
			Name:        "Real-time Updates",
			Description: "Get the latest posts as they're published",
			Icon:        "⚡",
		},
	}

	response := HomeResponse{
		AppName:     "RSS Aggregator",
		Version:     "1.0.0",
		Description: "A modern RSS feed aggregator with clean API",
		Timestamp:   time.Now(),
		Stats:       stats,
		Features:    features,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetAbout handles GET /about - about page
func (h *HomeHandler) GetAbout(w http.ResponseWriter, r *http.Request) {
	about := map[string]interface{}{
		"name":        "RSS Aggregator",
		"version":     "1.0.0",
		"description": "A modern RSS feed aggregator built with Go",
		"author":      "Your Name",
		"repository":  "https://github.com/olehhuss/rssagg",
		"license":     "MIT",
		"tech_stack": []string{
			"Go",
			"PostgreSQL",
			"Chi Router",
			"SQLC",
		},
		"features": []string{
			"RSS Feed Management",
			"User Authentication via API Keys",
			"Real-time Feed Updates",
			"RESTful API",
			"Clean Architecture",
		},
	}

	respondWithJSON(w, http.StatusOK, about)
}

// GetAPIInfo handles GET /api - API documentation info
func (h *HomeHandler) GetAPIInfo(w http.ResponseWriter, r *http.Request) {
	apiInfo := map[string]interface{}{
		"api_version": "v1",
		"base_url":    "/v1",
		"description": "RSS Aggregator REST API",
		"endpoints": map[string]interface{}{
			"health": map[string]string{
				"method":      "GET",
				"path":        "/v1/healthz",
				"description": "Health check endpoint",
				"auth":        "none",
			},
			"users": map[string]interface{}{
				"create": map[string]string{
					"method":      "POST",
					"path":        "/v1/users",
					"description": "Create a new user",
					"auth":        "none",
				},
				"me": map[string]string{
					"method":      "GET",
					"path":        "/v1/users/me",
					"description": "Get current user info",
					"auth":        "API Key required",
				},
			},
			"feeds": map[string]interface{}{
				"create": map[string]string{
					"method":      "POST",
					"path":        "/v1/feeds",
					"description": "Create a new feed",
					"auth":        "API Key required",
				},
				"list": map[string]string{
					"method":      "GET",
					"path":        "/v1/feeds",
					"description": "Get user's feeds",
					"auth":        "API Key required",
				},
				"all": map[string]string{
					"method":      "GET",
					"path":        "/v1/feeds/all",
					"description": "Get all feeds (admin)",
					"auth":        "API Key required",
				},
			},
		},
		"authentication": map[string]string{
			"type":        "API Key",
			"header":      "Authorization",
			"format":      "ApiKey YOUR_API_KEY_HERE",
			"description": "Include your API key in the Authorization header",
		},
		"examples": map[string]interface{}{
			"create_user": map[string]interface{}{
				"method": "POST",
				"url":    "/v1/users",
				"headers": map[string]string{
					"Content-Type": "application/json",
				},
				"body": map[string]string{
					"name": "John Doe",
				},
			},
			"get_feeds": map[string]interface{}{
				"method": "GET",
				"url":    "/v1/feeds",
				"headers": map[string]string{
					"Authorization": "ApiKey YOUR_API_KEY_HERE",
				},
			},
		},
	}

	respondWithJSON(w, http.StatusOK, apiInfo)
}
