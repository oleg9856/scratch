package handler

import (
	"encoding/json"
	"net/http"

	"github.com/olehhuss/rssagg/internal/usecase"
)

// FeedHandler handles HTTP requests for feeds
type FeedHandler struct {
	feedService *usecase.FeedService
}

// NewFeedHandler creates a new feed handler
func NewFeedHandler(feedService *usecase.FeedService) FeedHandlerInterface {
	return &FeedHandler{
		feedService: feedService,
	}
}

// CreateFeed handles POST /feeds
func (h *FeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request) {
	user, ok := getUserFromContext(r.Context())
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "User not found in context")
		return
	}

	var req usecase.CreateFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	// Set user ID from context
	req.UserID = user.ID

	feed, err := h.feedService.CreateFeed(r.Context(), req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, feed)
}

// GetUserFeeds handles GET /feeds
func (h *FeedHandler) GetUserFeeds(w http.ResponseWriter, r *http.Request) {
	user, ok := getUserFromContext(r.Context())
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "User not found in context")
		return
	}

	// TODO: Implement GetUserFeeds in service
	// feeds, err := h.feedService.GetUserFeeds(r.Context(), user.ID)
	// if err != nil {
	//     respondWithError(w, http.StatusInternalServerError, err.Error())
	//     return
	// }

	// Placeholder response
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "GetUserFeeds not implemented yet",
		"user_id": user.ID.String(),
	})
}

// GetAllFeeds handles GET /feeds/all (admin endpoint)
func (h *FeedHandler) GetAllFeeds(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement GetAllFeeds in service
	// feeds, err := h.feedService.GetAllFeeds(r.Context())
	// if err != nil {
	//     respondWithError(w, http.StatusInternalServerError, err.Error())
	//     return
	// }

	// Placeholder response
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "GetAllFeeds not implemented yet",
	})
}
