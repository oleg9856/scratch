package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

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
