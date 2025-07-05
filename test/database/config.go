package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/olehhuss/rssagg/internal/database"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
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

func GetOrCreateTestContainer() *TestContainer {
	once.Do(func() {
		instance = &TestContainer{
			Config: DefaultDatabaseConfig(),
		}
	})
	return instance
}

func (tc *TestContainer) Start(ctx context.Context) error {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	pgContainer, err := postgres.Run(ctx,
		tc.Config.Image,
		postgres.WithDatabase(tc.Config.DatabaseName),
		postgres.WithUsername(tc.Config.Username),
		postgres.WithPassword(tc.Config.Password),
		testcontainers.WithWaitStrategy(tc.Config.WaitStrategy),
	)
	if err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		pgContainer.Terminate(ctx)
		return fmt.Errorf("failed to get connection string: %w", err)
	}

	tc.Container = pgContainer
	tc.ConnString = connStr
	tc.isRunning = true

	return nil
}

func (tc *TestContainer) Stop(ctx context.Context) error {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	if !tc.isRunning || tc.Container == nil {
		return nil
	}

	err := tc.Container.Terminate(ctx)
	tc.isRunning = false
	tc.Container = nil
	tc.ConnString = ""

	return err
}

func (tc *TestContainer) GetDB() (*sql.DB, error) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	if !tc.isRunning {
		return nil, fmt.Errorf("container is not running")
	}

	db, err := sql.Open("postgres", tc.ConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// IsRunning перевіряє чи запущений контейнер
func (tc *TestContainer) IsRunning() bool {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	return tc.isRunning
}

func StartTestContainer() error {
	ctx := context.Background()
	testContainer := GetOrCreateTestContainer()
	return testContainer.Start(ctx)
}

func SetupTestDatabase(t *testing.T) (*database.Queries, func()) {
	ctx := context.Background()

	// Отримуємо singleton контейнер
	tc := GetOrCreateTestContainer()

	// Запускаємо контейнер
	err := tc.Start(ctx)
	require.NoError(t, err, "Failed to start test container")

	// Створюємо підключення
	db, err := tc.GetDB()
	require.NoError(t, err, "Failed to get database connection")

	// Створюємо таблиці
	err = createTestTables(db)
	require.NoError(t, err, "Failed to create test tables")

	queries := database.New(db)

	// Cleanup функція
	cleanup := func() {
		cleanupTestData(db) // Очищуємо дані
		db.Close()
	}

	return queries, cleanup
}

// SetupIsolatedTestDatabase створює окремий контейнер для тесту
func SetupIsolatedTestDatabase(t *testing.T) (*database.Queries, func()) {
	ctx := context.Background()

	config := DefaultDatabaseConfig()

	// Створюємо окремий контейнер
	pgContainer, err := postgres.Run(ctx,
		config.Image,
		postgres.WithDatabase(config.DatabaseName),
		postgres.WithUsername(config.Username),
		postgres.WithPassword(config.Password),
		testcontainers.WithWaitStrategy(config.WaitStrategy),
	)
	require.NoError(t, err)

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	require.NoError(t, db.Ping())

	err = createTestTables(db)
	require.NoError(t, err)

	queries := database.New(db)

	// Cleanup функція (видаляє весь контейнер)
	cleanup := func() {
		db.Close()
		pgContainer.Terminate(ctx)
	}

	return queries, cleanup
}

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

func cleanupTestData(db *sql.DB) {
	db.Exec("DELETE FROM feeds")
	db.Exec("DELETE FROM users")
}
