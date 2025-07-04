package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/domain"
	"github.com/olehhuss/rssagg/internal/interfaces"
)

//go:generate mockgen -source=user_service.go -destination=../../test/usecase/user_service_test.go -package=usecase interfaces.UserRepository

// UserService handles user-related business logic
type UserService struct {
	userRepo interfaces.UserRepository // ← Тепер інтерфейс!
}

// NewUserService creates a new user service
func NewUserService(userRepo interfaces.UserRepository) interfaces.UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user with generated API key
func (s *UserService) CreateUser(ctx context.Context, req interfaces.CreateUserRequest) (*domain.User, error) {
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

// GetUserByID retrieves user by ID
func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return user, nil
}

// GetUserCount returns the total number of users
func (s *UserService) GetUserCount(ctx context.Context) (int, error) {
	count, err := s.userRepo.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get user count: %w", err)
	}

	return count, nil
}

// generateAPIKey generates a random API key
func (s *UserService) generateAPIKey() string {
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	hash := sha256.Sum256(randomBytes)
	return hex.EncodeToString(hash[:])
}
