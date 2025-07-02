# RSS Aggregator (RSSAGG)

🗞️ **RSS Aggregator** - a web service for collecting, processing, and aggregating RSS feeds, built in Go using clean architecture.

## ✨ Features

- 📡 **RSS Parsing** - automatic downloading and parsing of RSS feeds
- 👥 **User Management** - registration and authentication via API keys
- 📰 **Feed Subscriptions** - users can subscribe to RSS feeds
- 🔄 **Automatic Updates** - periodic content updates from feeds
- 🏗️ **Clean Architecture** - separation into layers (domain, usecase, repository, handler)
- 🔧 **Builder Pattern** - convenient creation of domain objects
- 🗄️ **PostgreSQL** - reliable database with SQLC for type-safe queries


## 🚀 Quick Start

### Prerequisites
- Go 1.23+
- PostgreSQL 12+
- Make (optional)

### 1. Clone Repository
```bash
git clone <repository-url>
cd rssagg
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Database Setup
```bash
# Create PostgreSQL database
createdb rssagg

# Copy environment file
cp .env.example .env

# Edit .env file
nano .env
```

### 5. Run Migrations
```bash
# Install goose (if not installed)
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose -dir sql/schema postgres $DB_URL up
```

### 6. Build and Run
```bash
# Build
go build -o rssagg cmd/server/main.go

# Run
./rssagg
```

Server will start on `http://localhost:8082`

## 📚 API Documentation

### Authentication
All protected endpoints require header:
```
Authorization: ApiKey YOUR_API_KEY
```

### Endpoints

#### 👤 Users

**Create User**
```http
POST /v1/user
Content-Type: application/json

{
  "name": "John Doe"
}
```

**Get Current User** (protected)
```http
GET /v1/user/me
Authorization: ApiKey YOUR_API_KEY
```

#### 📡 Feeds

**Create Feed** (protected)
```http
POST /v1/feed
Authorization: ApiKey YOUR_API_KEY
Content-Type: application/json

{
  "name": "Tech Blog",
  "url": "https://example.com/rss"
}
```

**Get User Feeds** (protected)
```http
GET /v1/user-feeds
Authorization: ApiKey YOUR_API_KEY
```

#### 📰 Posts

**Get User Posts** (protected)
```http
GET /v1/posts?limit=10
Authorization: ApiKey YOUR_API_KEY
```

#### 🔄 Subscriptions

**Follow Feed** (protected)
```http
POST /v1/feeds/{feed_id}/follow
Authorization: ApiKey YOUR_API_KEY
```

**Unfollow Feed** (protected)
```http
DELETE /v1/feeds/{feed_id}/follow
Authorization: ApiKey YOUR_API_KEY
```

### 🔧 Health Check
```http
GET /v1/healthz
```

## 🗄️ Database

### Schema

#### Users
- `id` (UUID, PK)
- `name` (VARCHAR)
- `api_key` (VARCHAR, UNIQUE)
- `created_at`, `updated_at` (TIMESTAMP)

#### Feeds
- `id` (UUID, PK)
- `name` (VARCHAR)
- `url` (TEXT, UNIQUE)
- `user_id` (UUID, FK)
- `last_fetched_at` (TIMESTAMP, nullable)
- `created_at`, `updated_at` (TIMESTAMP)

#### Posts
- `id` (UUID, PK)
- `title` (TEXT)
- `description` (TEXT, nullable)
- `url` (TEXT, UNIQUE)
- `feed_id` (UUID, FK)
- `published_at` (TIMESTAMP, nullable)
- `created_at` (TIMESTAMP)

#### Feed Follows
- `id` (UUID, PK)
- `user_id` (UUID, FK)
- `feed_id` (UUID, FK)
- `created_at`, `updated_at` (TIMESTAMP)
- UNIQUE(user_id, feed_id)

## 🛠️ Development

### Generate SQLC Code
```bash
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Generate code
sqlc generate
```

### Adding New Migrations
```bash
# Create new migration
goose -dir sql/schema create migration_name sql

# Apply migration
goose -dir sql/schema postgres $DB_URL up
```

### Running Tests
```bash
# All tests
go test ./...

# Tests with coverage
go test -cover ./...

# Detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🏗️ Architectural Principles

### Domain Layer
- Clean business entities without dependencies
- Builder pattern for convenient object creation
- Nullable fields through pointer types

### Use Case Layer
- Business logic and orchestration
- Interfaces for repository layer
- Business rules validation

### Repository Layer
- Data access abstraction
- Conversion between domain and database models
- Type-safe queries through SQLC

### Handler Layer
- HTTP request/response processing
- Authentication and authorization
- JSON serialization/deserialization

