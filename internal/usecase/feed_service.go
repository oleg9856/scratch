package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
	"github.com/olehhuss/rssagg/internal/interfaces"
)

//go:generate mockgen -destination=../../test/mocks/feed_service_mock.go -package=mocks github.com/olehhuss/rssagg/internal/interfaces FeedServiceInterface,FeedRepository,UserRepository

// FeedService handles feed-related business logic
type FeedService struct {
	feedRepo interfaces.FeedRepository // ← Інтерфейс
	userRepo interfaces.UserRepository // ← Інтерфейс
}

// NewFeedService creates a new feed service
func NewFeedService(feedRepo interfaces.FeedRepository, userRepo interfaces.UserRepository) interfaces.FeedServiceInterface {
	return &FeedService{
		feedRepo: feedRepo,
		userRepo: userRepo,
	}
}

func (s *FeedService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Feed, error) {
	feeds, err := s.feedRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user feeds: %w", err)
	}

	return feeds, nil
}

// CreateFeed creates a new feed
func (s *FeedService) CreateFeed(ctx context.Context, req interfaces.CreateFeedRequest, userID uuid.UUID) (*domain.Feed, error) {
	// Валідація
	if req.Name == "" {
		return nil, fmt.Errorf("feed name is required")
	}
	if req.URL == "" {
		return nil, fmt.Errorf("feed URL is required")
	}

	// Перевірка існування користувача
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Створення фіду
	feed := domain.NewFeedBuilder().
		WithName(req.Name).
		WithURL(req.URL).
		WithUserID(userID).
		Build()

	// Збереження в базу
	if err := s.feedRepo.Create(ctx, feed); err != nil {
		return nil, fmt.Errorf("failed to create feed: %w", err)
	}

	return feed, nil
}

// GetAllFeeds retrieves all feeds
func (s *FeedService) GetAllFeeds(ctx context.Context) ([]*domain.Feed, error) {
	return s.feedRepo.GetAll(ctx)
}

// GetFeedCount returns the total number of feeds
func (s *FeedService) GetFeedCount(ctx context.Context) (int, error) {
	count, err := s.feedRepo.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get feed count: %w", err)
	}

	return count, nil
}
