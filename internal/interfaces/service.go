package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*domain.User, error)
	GetUserByAPIKey(ctx context.Context, apiKey string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetUserCount(ctx context.Context) (int, error)
}

type FeedServiceInterface interface {
	CreateFeed(ctx context.Context, req CreateFeedRequest, userID uuid.UUID) (*domain.Feed, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Feed, error)
	GetAllFeeds(ctx context.Context) ([]*domain.Feed, error)
	GetFeedCount(ctx context.Context) (int, error)
}

type PostServiceInterface interface {
	GetUserPosts(ctx context.Context, userID uuid.UUID, limit int) ([]*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) error
}

type FeedFollowServiceInterface interface {
	FollowFeed(ctx context.Context, userID, feedID uuid.UUID) (*domain.FeedFollow, error)
	UnfollowFeed(ctx context.Context, userID, feedID uuid.UUID) error
	GetUserFollows(ctx context.Context, userID uuid.UUID) ([]*domain.FeedFollow, error)
}

// Request structures
type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateFeedRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
