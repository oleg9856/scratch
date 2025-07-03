package handler

import "net/http"

// UserHandlerInterface defines the contract for user handlers
type UserHandlerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

// FeedHandlerInterface defines the contract for feed handlers
type FeedHandlerInterface interface {
	CreateFeed(w http.ResponseWriter, r *http.Request)
	GetUserFeeds(w http.ResponseWriter, r *http.Request)
	GetAllFeeds(w http.ResponseWriter, r *http.Request)
}

// PostHandlerInterface defines the contract for post handlers
type PostHandlerInterface interface {
	GetUserPosts(w http.ResponseWriter, r *http.Request)
}

// FeedFollowHandlerInterface defines the contract for feed follow handlers
type FeedFollowHandlerInterface interface {
	FollowFeed(w http.ResponseWriter, r *http.Request)
	UnfollowFeed(w http.ResponseWriter, r *http.Request)
	GetUserFollows(w http.ResponseWriter, r *http.Request)
}

// AuthMiddlewareInterface defines the contract for auth middleware
type AuthMiddlewareInterface interface {
	Authenticate(next http.HandlerFunc) http.HandlerFunc
}

// HomeHandlerInterface defines the contract for home/landing page handlers
type HomeHandlerInterface interface {
	GetHome(w http.ResponseWriter, r *http.Request)
	GetAbout(w http.ResponseWriter, r *http.Request)
	GetAPIInfo(w http.ResponseWriter, r *http.Request)
}
