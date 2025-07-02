package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/database"
	"github.com/olehhuss/rssagg/internal/domain"
	"github.com/olehhuss/rssagg/internal/usecase"
)

// FeedRepository implements the feed repository using PostgreSQL
type FeedRepository struct {
	queries *database.Queries
}

// Create implements usecase.FeedRepository.
func (f *FeedRepository) Create(ctx context.Context, feed *domain.Feed) error {
	params := database.CreateFeedParams{
		ID:        feed.ID,
		CreatedAt: sql.NullTime{Time: feed.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: feed.UpdatedAt, Valid: true},
		Name:      feed.Name,
		Url:       feed.URL, // Note: database uses 'Url', domain uses 'URL'
		UserID:    feed.UserID,
	}

	dbFeed, err := f.queries.CreateFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	// Update domain feed with database values (like LastFetchedAt)
	if dbFeed.LastFetchedAt.Valid {
		feed.LastFetchedAt = &dbFeed.LastFetchedAt.Time
	}

	return nil
}

// GetAll implements usecase.FeedRepository.
func (f *FeedRepository) GetAll(ctx context.Context) ([]*domain.Feed, error) {
	dbFeeds, err := f.queries.GetAllFeeds(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all feeds: %w", err)
	}

	feeds := make([]*domain.Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = f.convertToDomain(dbFeed)
	}

	return feeds, nil
}

// GetByUserID implements usecase.FeedRepository.
func (f *FeedRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Feed, error) {
	dbFeeds, err := f.queries.GetFeedsByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get feeds by user: %w", err)
	}

	feeds := make([]*domain.Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = f.convertToDomain(dbFeed)
	}

	return feeds, nil
}

// MarkAsFetched implements usecase.FeedRepository.
func (f *FeedRepository) MarkAsFetched(ctx context.Context, feedID uuid.UUID) error {
	err := f.queries.MarkFeedAsFetched(ctx, feedID)
	if err != nil {
		return fmt.Errorf("failed to mark feed as fetched: %w", err)
	}
	return nil
}

// convertToDomain converts database Feed to domain Feed
func (f *FeedRepository) convertToDomain(dbFeed database.Feed) *domain.Feed {
	feed := &domain.Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url, // Convert Url to URL
		UserID:    dbFeed.UserID,
		CreatedAt: f.nullTimeToTime(dbFeed.CreatedAt),
		UpdatedAt: f.nullTimeToTime(dbFeed.UpdatedAt),
	}

	// Handle nullable LastFetchedAt
	if dbFeed.LastFetchedAt.Valid {
		feed.LastFetchedAt = &dbFeed.LastFetchedAt.Time
	}

	return feed
}

// nullTimeToTime converts sql.NullTime to time.Time
func (f *FeedRepository) nullTimeToTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func NewFeedRepository(queries *database.Queries) usecase.FeedRepository {
	return &FeedRepository{
		queries: queries,
	}
}
