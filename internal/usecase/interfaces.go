package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

// UserRepository defines the contract for user data access
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*domain.User, error)
}

// FeedRepository defines the contract for feed data access
type FeedRepository interface {
	Create(ctx context.Context, feed *domain.Feed) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Feed, error)
	GetAll(ctx context.Context) ([]*domain.Feed, error)
	MarkAsFetched(ctx context.Context, feedID uuid.UUID) error
}

// PostRepository defines the contract for post data access
type PostRepository interface {
	Create(ctx context.Context, post *domain.Post) error
	GetByUserID(ctx context.Context, userID uuid.UUID, limit int) ([]*domain.Post, error)
	Exists(ctx context.Context, url string) (bool, error)
}

// FeedFollowRepository defines the contract for feed follows
type FeedFollowRepository interface {
	Create(ctx context.Context, follow *domain.FeedFollow) error
	Delete(ctx context.Context, userID, feedID uuid.UUID) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.FeedFollow, error)
}
