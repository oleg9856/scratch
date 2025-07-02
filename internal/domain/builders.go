package domain

import (
	"time"

	"github.com/google/uuid"
)

// UserBuilder helps build User entities
type UserBuilder struct {
	user *User
}

// NewUserBuilder creates a new UserBuilder
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{
		user: &User{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// WithID sets the user ID
func (b *UserBuilder) WithID(id uuid.UUID) *UserBuilder {
	b.user.ID = id
	return b
}

// WithName sets the user name
func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

// WithAPIKey sets the API key
func (b *UserBuilder) WithAPIKey(apiKey string) *UserBuilder {
	b.user.APIKey = apiKey
	return b
}

// WithCreatedAt sets the creation time
func (b *UserBuilder) WithCreatedAt(createdAt time.Time) *UserBuilder {
	b.user.CreatedAt = createdAt
	return b
}

// WithUpdatedAt sets the update time
func (b *UserBuilder) WithUpdatedAt(updatedAt time.Time) *UserBuilder {
	b.user.UpdatedAt = updatedAt
	return b
}

// Build returns the built User
func (b *UserBuilder) Build() *User {
	return b.user
}

// FeedBuilder helps build Feed entities
type FeedBuilder struct {
	feed *Feed
}

// NewFeedBuilder creates a new FeedBuilder
func NewFeedBuilder() *FeedBuilder {
	return &FeedBuilder{
		feed: &Feed{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// WithID sets the feed ID
func (b *FeedBuilder) WithID(id uuid.UUID) *FeedBuilder {
	b.feed.ID = id
	return b
}

// WithName sets the feed name
func (b *FeedBuilder) WithName(name string) *FeedBuilder {
	b.feed.Name = name
	return b
}

// WithURL sets the feed URL
func (b *FeedBuilder) WithURL(url string) *FeedBuilder {
	b.feed.URL = url
	return b
}

// WithUserID sets the user ID
func (b *FeedBuilder) WithUserID(userID uuid.UUID) *FeedBuilder {
	b.feed.UserID = userID
	return b
}

// WithLastFetchedAt sets the last fetched time
func (b *FeedBuilder) WithLastFetchedAt(lastFetchedAt *time.Time) *FeedBuilder {
	b.feed.LastFetchedAt = lastFetchedAt
	return b
}

// WithCreatedAt sets the creation time
func (b *FeedBuilder) WithCreatedAt(createdAt time.Time) *FeedBuilder {
	b.feed.CreatedAt = createdAt
	return b
}

// WithUpdatedAt sets the update time
func (b *FeedBuilder) WithUpdatedAt(updatedAt time.Time) *FeedBuilder {
	b.feed.UpdatedAt = updatedAt
	return b
}

// Build returns the built Feed
func (b *FeedBuilder) Build() *Feed {
	return b.feed
}

// PostBuilder helps build Post entities
type PostBuilder struct {
	post *Post
}

// NewPostBuilder creates a new PostBuilder
func NewPostBuilder() *PostBuilder {
	return &PostBuilder{
		post: &Post{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
		},
	}
}

// WithID sets the post ID
func (b *PostBuilder) WithID(id uuid.UUID) *PostBuilder {
	b.post.ID = id
	return b
}

// WithTitle sets the post title
func (b *PostBuilder) WithTitle(title string) *PostBuilder {
	b.post.Title = title
	return b
}

// WithDescription sets the post description
func (b *PostBuilder) WithDescription(description *string) *PostBuilder {
	b.post.Description = description
	return b
}

// WithDescriptionString sets the post description from string
func (b *PostBuilder) WithDescriptionString(description string) *PostBuilder {
	b.post.Description = &description
	return b
}

// WithURL sets the post URL
func (b *PostBuilder) WithURL(url string) *PostBuilder {
	b.post.URL = url
	return b
}

// WithFeedID sets the feed ID
func (b *PostBuilder) WithFeedID(feedID uuid.UUID) *PostBuilder {
	b.post.FeedID = feedID
	return b
}

// WithPublishedAt sets the published time
func (b *PostBuilder) WithPublishedAt(publishedAt *time.Time) *PostBuilder {
	b.post.PublishedAt = publishedAt
	return b
}

// WithPublishedAtTime sets the published time from time.Time
func (b *PostBuilder) WithPublishedAtTime(publishedAt time.Time) *PostBuilder {
	b.post.PublishedAt = &publishedAt
	return b
}

// WithCreatedAt sets the creation time
func (b *PostBuilder) WithCreatedAt(createdAt time.Time) *PostBuilder {
	b.post.CreatedAt = createdAt
	return b
}

// Build returns the built Post
func (b *PostBuilder) Build() *Post {
	return b.post
}

// FeedFollowBuilder helps build FeedFollow entities
type FeedFollowBuilder struct {
	feedFollow *FeedFollow
}

// NewFeedFollowBuilder creates a new FeedFollowBuilder
func NewFeedFollowBuilder() *FeedFollowBuilder {
	return &FeedFollowBuilder{
		feedFollow: &FeedFollow{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// WithID sets the feed follow ID
func (b *FeedFollowBuilder) WithID(id uuid.UUID) *FeedFollowBuilder {
	b.feedFollow.ID = id
	return b
}

// WithUserID sets the user ID
func (b *FeedFollowBuilder) WithUserID(userID uuid.UUID) *FeedFollowBuilder {
	b.feedFollow.UserID = userID
	return b
}

// WithFeedID sets the feed ID
func (b *FeedFollowBuilder) WithFeedID(feedID uuid.UUID) *FeedFollowBuilder {
	b.feedFollow.FeedID = feedID
	return b
}

// WithCreatedAt sets the creation time
func (b *FeedFollowBuilder) WithCreatedAt(createdAt time.Time) *FeedFollowBuilder {
	b.feedFollow.CreatedAt = createdAt
	return b
}

// WithUpdatedAt sets the update time
func (b *FeedFollowBuilder) WithUpdatedAt(updatedAt time.Time) *FeedFollowBuilder {
	b.feedFollow.UpdatedAt = updatedAt
	return b
}

// Build returns the built FeedFollow
func (b *FeedFollowBuilder) Build() *FeedFollow {
	return b.feedFollow
}
