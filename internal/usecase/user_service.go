package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/olehhuss/rssagg/internal/domain"
)

// UserService handles user-related business logic
type UserService struct {
	userRepo UserRepository
}

// NewUserService creates a new user service
func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUserRequest represents the request to create a user
type CreateUserRequest struct {
	Name string `json:"name"`
}

// CreateUser creates a new user with generated API key
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*domain.User, error) {
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	user := domain.NewUserBuilder().
		WithName(req.Name).
		WithAPIKey(s.generateAPIKey()).
		Build()

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// GetUserByAPIKey retrieves user by API key
func (s *UserService) GetUserByAPIKey(ctx context.Context, apiKey string) (*domain.User, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	user, err := s.userRepo.GetByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return user, nil
}

// generateAPIKey generates a random API key
func (s *UserService) generateAPIKey() string {
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:])
}
