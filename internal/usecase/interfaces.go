package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

// Repository interfaces - визначаються в usecase, реалізуються в repository
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*domain.User, error)
	Count(ctx context.Context) (int, error)
}

type FeedRepository interface {
	Create(ctx context.Context, feed *domain.Feed) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Feed, error)
	GetAll(ctx context.Context) ([]*domain.Feed, error)
	MarkAsFetched(ctx context.Context, feedID uuid.UUID) error
	Count(ctx context.Context) (int, error)
}

type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByUserID(ctx context.Context, userID uuid.UUID, limit int) ([]*domain.Post, error)
	Exists(ctx context.Context, url string) (bool, error)
}

type FeedFollowRepository interface {
	Create(ctx context.Context, follow *domain.FeedFollow) error
	Delete(ctx context.Context, userID, feedID uuid.UUID) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.FeedFollow, error)
}

// Service interfaces - для handler шару
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
