package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
)

func TestUserBuilder(t *testing.T) {
	// Simple user
	user := domain.NewUserBuilder().
		WithName("John Doe").
		WithAPIKey("test-key").
		Build()

	if user.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got %s", user.Name)
	}

	if user.APIKey != "test-key" {
		t.Errorf("Expected API key 'test-key', got %s", user.APIKey)
	}

	// User with custom timestamps
	customTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	userWithTime := domain.NewUserBuilder().
		WithName("Jane Doe").
		WithCreatedAt(customTime).
		WithUpdatedAt(customTime).
		Build()

	if !userWithTime.CreatedAt.Equal(customTime) {
		t.Errorf("Expected custom creation time")
	}
}

func TestFeedBuilder(t *testing.T) {
	userID := uuid.New()

	// Feed without last fetched time
	feed := domain.NewFeedBuilder().
		WithName("Tech News").
		WithURL("https://example.com/rss").
		WithUserID(userID).
		Build()

	if feed.Name != "Tech News" {
		t.Errorf("Expected name 'Tech News', got %s", feed.Name)
	}

	if feed.LastFetchedAt != nil {
		t.Errorf("Expected LastFetchedAt to be nil")
	}

	// Feed with last fetched time
	fetchTime := time.Now()
	feedWithFetch := domain.NewFeedBuilder().
		WithName("News Feed").
		WithURL("https://news.com/rss").
		WithUserID(userID).
		WithLastFetchedAt(&fetchTime).
		Build()

	if feedWithFetch.LastFetchedAt == nil {
		t.Errorf("Expected LastFetchedAt to be set")
	}
}

func TestPostBuilder(t *testing.T) {
	feedID := uuid.New()

	// Post with description
	post := domain.NewPostBuilder().
		WithTitle("Test Article").
		WithDescriptionString("This is a test article").
		WithURL("https://example.com/article").
		WithFeedID(feedID).
		WithPublishedAtTime(time.Now()).
		Build()

	if post.Title != "Test Article" {
		t.Errorf("Expected title 'Test Article', got %s", post.Title)
	}

	if post.Description == nil || *post.Description != "This is a test article" {
		t.Errorf("Expected description to be set")
	}

	// Post without description
	postNoDesc := domain.NewPostBuilder().
		WithTitle("Another Article").
		WithURL("https://example.com/another").
		WithFeedID(feedID).
		Build()

	if postNoDesc.Description != nil {
		t.Errorf("Expected description to be nil")
	}
}

func TestFeedFollowBuilder(t *testing.T) {
	userID := uuid.New()
	feedID := uuid.New()

	follow := domain.NewFeedFollowBuilder().
		WithUserID(userID).
		WithFeedID(feedID).
		Build()

	if follow.UserID != userID {
		t.Errorf("Expected user ID to match")
	}

	if follow.FeedID != feedID {
		t.Errorf("Expected feed ID to match")
	}
}

// Example of creating complex objects for testing
func ExampleBuildersUsage() {
	// Create a user
	user := domain.NewUserBuilder().
		WithName("Alice Smith").
		WithAPIKey("alice-secret-key").
		Build()

	// Create a feed for the user
	feed := domain.NewFeedBuilder().
		WithName("Alice's Tech Blog").
		WithURL("https://alice.dev/rss").
		WithUserID(user.ID).
		Build()

	// Create a post in the feed
	post := domain.NewPostBuilder().
		WithTitle("Understanding Go Builders").
		WithDescriptionString("Learn how to implement the builder pattern in Go").
		WithURL("https://alice.dev/go-builders").
		WithFeedID(feed.ID).
		WithPublishedAtTime(time.Now()).
		Build()

	// Create a feed follow relationship
	follow := domain.NewFeedFollowBuilder().
		WithUserID(user.ID).
		WithFeedID(feed.ID).
		Build()

	_ = user
	_ = feed
	_ = post
	_ = follow
}
