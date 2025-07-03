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

// UserRepository implements usecase.UserRepository interface
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
	params := r.domainUserToCreateParams(user)

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
	dbUser, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return r.dbUserToDomain(dbUser), nil
}

// GetByAPIKey retrieves a user by API key
func (r *UserRepository) GetByAPIKey(ctx context.Context, apiKey string) (*domain.User, error) {
	dbUser, err := r.queries.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by API key: %w", err)
	}

	return r.dbUserToDomain(dbUser), nil
}

// Count returns the total number of users
func (r *UserRepository) Count(ctx context.Context) (int, error) {
	// Note: You'll need to add this query to users.sql if it doesn't exist
	count, err := r.queries.GetUserCount(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to count users: %w", err)
	}

	return int(count), nil
}

// dbUserToDomain converts a database user model to a domain user model
func (r *UserRepository) dbUserToDomain(dbUser database.User) *domain.User {
	return &domain.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt.Time,
		UpdatedAt: dbUser.UpdatedAt.Time,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

// domainUserToCreateParams converts a domain user to database create parameters
func (r *UserRepository) domainUserToCreateParams(user *domain.User) database.CreateUserParams {
	return database.CreateUserParams{
		ID:        user.ID,
		CreatedAt: sql.NullTime{Time: user.CreatedAt, Valid: true},
		UpdatedAt: sql.NullTime{Time: user.UpdatedAt, Valid: true},
		Name:      user.Name,
	}
}
