package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/olehhuss/rssagg/internal/database"
	"github.com/olehhuss/rssagg/internal/domain"
	"github.com/olehhuss/rssagg/internal/usecase"
)

// UserRepository implements the user repository using PostgreSQL
type UserRepository struct {
	queries *database.Queries
}

// NewUserRepository creates a new user repository
func NewUserRepository(queries *database.Queries) usecase.UserRepository {
	return &UserRepository{
		queries: queries,
	}
}

// Create creates a new user in the database
func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	params := database.CreateUserParams{
		ID:        user.ID,
		CreatedAt: sql.NullTime{Time: user.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt, Valid: true},
		Name:      user.Name,
	}

	dbUser, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	// Update the domain user with the API key from database
	user.APIKey = dbUser.ApiKey
	return nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	// Note: You'll need to add this query to users.sql
	// For now, this is a placeholder
	return nil, fmt.Errorf("GetByID not implemented yet")
}

// GetByAPIKey retrieves a user by API key
func (r *UserRepository) GetByAPIKey(ctx context.Context, apiKey string) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by API key: %w", err)
	}

	domainUser := &domain.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}

	return domainUser, nil
}
