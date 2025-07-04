package database

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DatabaseConfig struct {
	Image        string
	DatabaseName string
	Username     string
	Password     string
	WaitStrategy wait.Strategy
}

type TestContainer struct {
	Container  *postgres.PostgresContainer
	ConnString string
	Config     *DatabaseConfig
	mu         sync.RWMutex
	isRunning  bool
}

var (
	instance *TestContainer
	once     sync.Once
)

func DefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Image:        "postgres:15-alpine",
		DatabaseName: "testdb",
		Username:     "testuser",
		Password:     "testpass",
		WaitStrategy: wait.ForLog("database system is ready to accept connections").
			WithOccurrence(2).
			WithStartupTimeout(30 * time.Second),
	}
}

func getOrCreateTestContainer() *TestContainer {
	once.Do(func() {
		instance = &TestContainer{
			Config: DefaultDatabaseConfig(),
		}
	})
	return instance
}

//nolint:unused
func createTestTables(db *sql.DB) error {
	schema := `
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
        name TEXT NOT NULL,
        api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
            encode(sha256(random()::text::bytea), 'hex')
        )
    );

    CREATE TABLE IF NOT EXISTS feeds (
        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
        name TEXT NOT NULL,
        url TEXT UNIQUE NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        last_fetched_at TIMESTAMP
    );`

	_, err := db.Exec(schema)
	return err
}
