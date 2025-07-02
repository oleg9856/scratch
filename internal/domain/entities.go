package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the domain
type User struct {
	ID        uuid.UUID
	Name      string
	APIKey    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Feed represents an RSS feed entity
type Feed struct {
	ID            uuid.UUID
	Name          string
	URL           string
	UserID        uuid.UUID
	LastFetchedAt *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Post represents a post from RSS feed
type Post struct {
	ID          uuid.UUID
	Title       string
	Description *string
	URL         string
	FeedID      uuid.UUID
	PublishedAt *time.Time
	CreatedAt   time.Time
}

// FeedFollow represents user subscription to feed
type FeedFollow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
