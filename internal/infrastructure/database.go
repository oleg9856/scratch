package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URL string
}

// NewDatabaseConfig creates database config from environment
func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		URL: os.Getenv("DB_URL"),
	}
}

// NewDatabase creates a new database connection
func NewDatabase(config *DatabaseConfig) (*sql.DB, error) {
	if config.URL == "" {
		return nil, fmt.Errorf("DB_URL environment variable not set")
	}

	db, err := sql.Open("postgres", config.URL)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot ping database: %w", err)
	}

	return db, nil
}
