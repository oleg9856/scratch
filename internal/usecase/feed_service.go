package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

// FeedService handles feed-related business logic
type FeedService struct {
	feedRepo FeedRepository
	userRepo UserRepository
}

// NewFeedService creates a new feed service
func NewFeedService(feedRepo FeedRepository, userRepo UserRepository) *FeedService {
	return &FeedService{
		feedRepo: feedRepo,
		userRepo: userRepo,
	}
}

// CreateFeedRequest represents the request to create a feed
type CreateFeedRequest struct {
	Name   string    `json:"name"`
	URL    string    `json:"url"`
	UserID uuid.UUID `json:"user_id"`
}

// CreateFeed creates a new feed
func (s *FeedService) CreateFeed(ctx context.Context, req CreateFeedRequest) (*domain.Feed, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if req.URL == "" {
		return nil, fmt.Errorf("url is required")
	}

	// Validate user
	_, err := s.userRepo.GetByID(ctx, req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to identify user: %w", err)
	}

	feed := domain.NewFeedBuilder().
		WithName(req.Name).
		WithURL(req.URL).
		WithUserID(req.UserID).
		Build()

	err = s.feedRepo.Create(ctx, feed)
	if err != nil {
		return nil, fmt.Errorf("failed to create feed: %w", err)
	}

	return feed, nil
}
