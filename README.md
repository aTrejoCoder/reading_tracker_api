# 📚 Reading Tracker API

> Enterprise-grade RESTful API for comprehensive reading management built with Go, MongoDB, and Clean Architecture principles.

[![Go Version](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8)](https://gin-gonic.com/)
[![Database](https://img.shields.io/badge/Database-MongoDB-47A248?logo=mongodb&logoColor=white)](https://www.mongodb.com/)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![API Docs](https://img.shields.io/badge/API-Swagger-85EA2D?logo=swagger)](http://localhost:8080/swagger/index.html)

---

## 🎯 Overview

**Reading Tracker API** is a production-ready backend service designed to solve the modern reader's challenge of managing diverse content formats. Whether tracking traditional books, manga collections, or custom documents, this API provides a unified platform with intelligent progress monitoring, secure multi-user support, and flexible organizational tools.

Built with **Clean Architecture** and **Go generics**, the system achieves enterprise-level code quality while maintaining high performance through Docker containerization and optimized MongoDB operations.

### Key Capabilities

- **Multi-Format Support**: Track books (with ISBN/metadata), manga (volumes/chapters), and custom documents
- **Progress Intelligence**: Historical reading records with timestamps, notes, and progress snapshots
- **Flexible Organization**: User-defined reading lists with custom categorization
- **Enterprise Security**: JWT authentication, bcrypt password hashing, and IP-based rate limiting (30 req/min)
- **Developer Experience**: Full Swagger documentation, type-safe operations, and containerized development

---

## 🏗️ Architecture

### Design Philosophy

The API follows **Clean Architecture** (aka Hexagonal Architecture) with strict layer separation ensuring testability, maintainability, and scalability:

```
┌─────────────────────────────────────────────────────────────┐
│                        HTTP Layer (Gin)                      │
│  ┌─────────────┐  ┌─────────────┐  ┌──────────────────┐    │
│  │ Rate Limiter│  │  JWT Auth   │  │  CORS/Security   │    │
│  └─────────────┘  └─────────────┘  └──────────────────┘    │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                    Controllers Layer                         │
│   AuthController │ UserController │ BookController │         │
│   ReadingController │ MangaController │ ListController       │
│  (HTTP handling, validation, response formatting)            │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                     Services Layer                           │
│   AuthService │ UserService │ BookService │ ReadingService   │
│  (Business logic, orchestration, DTO mapping)                │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                   Repository Layer                           │
│   Repository[T] (Generic) │ UserRepository (Specialized)     │
│  (Data access abstraction, MongoDB operations)               │
└───────────────────────────┬─────────────────────────────────┘
                            │
                      ┌─────▼─────┐
                      │  MongoDB  │
                      └───────────┘
```

### Core Design Patterns

- **Generic Repository Pattern**: Type-safe `Repository[T]` using Go 1.22 generics eliminates 70% of repository code duplication
- **Dependency Injection**: Constructor-based DI enabling isolated unit testing and loose coupling
- **DTO Pattern**: Separation between API contracts (DTOs) and domain models (Entities)
- **Middleware Chain**: Composable request processing for authentication, rate limiting, and logging
- **Service Layer**: Interface-driven business logic encapsulation for each domain

---

## 📁 Project Structure

```
reading-tracker-api/
│
├── main.go                          # Application entry point, dependency wiring
├── go.mod / go.sum                  # Go module dependencies
├── docker-compose.yml               # Multi-container orchestration
├── dockerfile                       # Multi-stage build configuration
│
├── controllers/                     # HTTP handlers (Gin)
│   ├── auth_controller.go          # Signup, Login endpoints
│   ├── user_controller.go          # User profile management
│   ├── book_controller.go          # Book CRUD operations
│   ├── manga_controller.go         # Manga catalog management
│   ├── reading_controller.go       # Reading session handling
│   ├── reading_list_controller.go  # Collection organization
│   ├── reading_record_controller.go # Progress tracking
│   └── custom_document_controller.go # Custom content management
│
├── services/                        # Business logic layer
│   ├── auth_service.go             # Authentication workflows
│   ├── user_service.go             # User business rules
│   ├── book_service.go             # Book cataloging logic
│   ├── manga_service.go            # Manga management
│   ├── reading_service.go          # Reading orchestration
│   ├── reading_list_service.go     # List organization
│   └── reading_record_service.go   # Progress calculation
│
├── repository/                      # Data access layer
│   ├── common_repository.go        # Generic Repository[T] implementation
│   ├── user_repository.go          # User-specific queries
│   ├── reading_repository.go       # Reading aggregations
│   ├── reading_list_repository.go  # List operations
│   └── custom_document_repository.go # Document queries
│
├── models/                          # Domain entities
│   ├── user.go                     # User, Profile entities
│   ├── reading.go                  # Reading, ReadingRecord, ReadingList
│   └── reading_types.go            # Book, Manga, CustomDocument
│
├── dtos/                            # Data Transfer Objects
│   ├── auth_dtos.go                # SignupDTO, LoginDTO
│   ├── user_dtos.go                # UserDTO, UserInsertDTO
│   ├── book_dtos.go                # BookDTO, BookInsertDTO
│   ├── manga_dtos.go               # MangaDTO, MangaInsertDTO
│   └── reading_dtos.go             # ReadingDTO, ReadingInsertDTO
│
├── mappers/                         # Entity ↔ DTO transformations
│   ├── auth_mapper.go
│   ├── user_mapper.go
│   ├── book_mapper.go
│   └── reading_mapper.go
│
├── middleware/                      # Cross-cutting concerns
│   ├── rate_limiter.go             # IP-based throttling (30/min)
│   └── token/
│       └── jwt.go                  # JWT generation & validation
│
├── routes/                          # Route definitions
│   ├── auth_routes.go              # /api/signup, /api/login
│   ├── user_routes.go              # /api/users/* endpoints
│   ├── reading_routes.go           # /api/readings/* endpoints
│   └── reading_documents_routes.go # /api/books/*, /api/mangas/*
│
├── utils/                           # Shared utilities
│   ├── api_response.go             # Standardized API responses
│   ├── errors.go                   # Custom error definitions
│   ├── hash_password.go            # bcrypt password handling
│   └── gin_helpers.go              # Request parsing utilities
│
├── database/                        # Database configuration
│   ├── db_setup.go                 # MongoDB connection setup
│   ├── db_collections.go           # Collection definitions
│   └── test/                       # Integration tests
│       ├── setup.go
│       ├── book_test.go
│       ├── manga_test.go
│       ├── reading_test.go
│       └── user_test.go
│
├── docs/                            # Auto-generated Swagger docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
└── mongo-init/                      # MongoDB initialization
    ├── app/
    │   └── init.js                 # Production DB setup
    └── test/
        └── init_test.js            # Test DB setup
```

### Layer Responsibilities

| Layer | Responsibility | Technology |
|-------|---------------|------------|
| **Controllers** | HTTP request/response handling, input validation, error formatting | Gin, Go Playground Validator |
| **Services** | Business logic, transaction orchestration, DTO-Entity mapping | Pure Go interfaces |
| **Repositories** | Data access abstraction, MongoDB queries, error translation | Go Generics, MongoDB Driver |
| **Middleware** | Authentication, rate limiting, CORS, logging | Gin middleware, JWT |
| **Models** | Domain entity definitions, BSON mapping | MongoDB BSON tags |

---

## 🚀 Quick Start

### Prerequisites

- **Go 1.22+** - [Install Go](https://go.dev/dl/)
- **Docker & Docker Compose** - [Install Docker](https://docs.docker.com/get-docker/)
- **MongoDB** (if running locally) - [Install MongoDB](https://www.mongodb.com/try/download/community)

### Installation & Running

#### Option 1: Docker Compose (Recommended)

```bash
# Clone the repository
git clone https://github.com/aTrejoCoder/reading_tracker_api.git
cd reading_tracker_api

# Start all services (app + MongoDB + test MongoDB)
docker-compose up --build

# API available at http://localhost:8080
# Swagger docs at http://localhost:8080/swagger/index.html
```

#### Option 2: Local Development

```bash
# Install dependencies
go mod download

# Set up environment variables
echo "JWT_SECRET=your-secret-key-here" > .env
echo "DATABASE_URL=mongodb://localhost:27017/reading_tracker" >> .env

# Run MongoDB (separate terminal)
mongod --dbpath ./data

# Run the application
go run main.go

# API available at http://localhost:8080
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./database/test/...

# Run specific test file
go test ./database/test/book_test.go
```

---

## 📊 Technical Highlights

### Performance & Scalability

- **Stateless Design**: JWT-based authentication enables horizontal scaling without session storage
- **Connection Pooling**: MongoDB driver manages connection pool for optimal resource utilization
- **Generic Repository**: Type-safe patterns reduce code paths and improve maintainability
- **Pagination**: Built-in pagination across all list endpoints preventing large data transfers
- **Rate Limiting**: 30 requests/minute per IP protecting against abuse and DDoS

### Security Features

- **JWT Authentication**: Token-based auth with 72-hour expiration and role-based claims
- **Password Security**: bcrypt hashing with configurable cost factor
- **Input Validation**: Comprehensive DTO validation preventing injection attacks
- **Environment Variables**: Secret management via .env files preventing credential leakage
- **Rate Limiting**: IP-based throttling protecting against brute force attacks

### Code Quality

- **Type Safety**: Go generics eliminating runtime type errors across repository layer
- **Clean Architecture**: Strict layer separation with dependency inversion
- **Interface-Driven**: All services defined by interfaces enabling easy mocking
- **Comprehensive Testing**: Integration tests for database operations
- **Swagger Documentation**: 100% API coverage with auto-generated documentation

---

## 🔧 Key Technologies

### Backend Stack

- **[Go 1.22](https://go.dev/)** - Modern, type-safe language with excellent performance
- **[Gin Web Framework](https://gin-gonic.com/)** - High-performance HTTP router with middleware support
- **[MongoDB Go Driver](https://www.mongodb.com/docs/drivers/go/current/)** - Official MongoDB driver with connection pooling
- **[JWT (golang-jwt/jwt)](https://github.com/golang-jwt/jwt)** - JSON Web Token implementation for authentication

### Infrastructure

- **[Docker](https://www.docker.com/)** - Multi-stage builds optimizing production image size (~25MB)
- **[Docker Compose](https://docs.docker.com/compose/)** - Three-service orchestration (app, production DB, test DB)
- **[MongoDB](https://www.mongodb.com/)** - Document-oriented database with flexible schemas

### Development Tools

- **[Swagger/OpenAPI](https://swagger.io/)** - API documentation with interactive testing interface
- **[Go Playground Validator](https://github.com/go-playground/validator)** - Comprehensive struct validation
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable management
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Secure password hashing

---

## 📖 API Documentation

### Interactive Swagger Documentation

Once the application is running, access the full API documentation at:

**http://localhost:8080/swagger/index.html**

### Sample Endpoints

#### Authentication
```http
POST /api/signup        # Create new user account
POST /api/login         # Authenticate and receive JWT token
```

#### Books & Manga
```http
GET    /api/books                # List all books (paginated)
POST   /api/books                # Create new book
GET    /api/books/{id}           # Get book by ID
GET    /api/books/isbn/{isbn}   # Get book by ISBN
PUT    /api/books/{id}           # Update book
DELETE /api/books/{id}           # Delete book

GET    /api/mangas               # List all manga (paginated)
POST   /api/mangas               # Create new manga
GET    /api/mangas/{id}          # Get manga by ID
```

#### Reading Management
```http
GET    /api/user-readings                # Get my readings
GET    /api/user-readings/by-status      # Filter by status
GET    /api/user-readings/by-type        # Filter by type
POST   /api/user-readings                # Start new reading
PUT    /api/user-readings/{id}           # Update reading
DELETE /api/user-readings/{id}           # Delete reading
```

#### Progress Tracking
```http
GET    /api/user-records/{id}   # Get reading progress history
POST   /api/user-records         # Add progress record
PUT    /api/user-records/{id}   # Update progress record
DELETE /api/user-records/{id}   # Delete progress record
```

#### Reading Lists
```http
GET    /api/user-lists           # Get my reading lists
POST   /api/user-lists           # Create new list
PUT    /api/user-lists/{id}      # Update list
DELETE /api/user-lists/{id}      # Delete list
```

---

## 🎨 Code Examples

### Generic Repository Pattern

```go
// Type-safe repository for any entity
type Repository[T any] struct {
    collection *mongo.Collection
}

func NewRepository[T any](collection *mongo.Collection) *Repository[T] {
    return &Repository[T]{collection: collection}
}

// Works with User, Book, Manga, Reading - no code duplication!
userRepo := repository.NewRepository[models.User](userCollection)
bookRepo := repository.NewRepository[models.Book](bookCollection)
```

### JWT Authentication

```go
type CustomClaims struct {
    ObjectId string   `json:"objectId"`
    Username string   `json:"username"`
    Email    string   `json:"email"`
    Roles    []string `json:"roles"`
    jwt.RegisteredClaims
}

func GenerateJWT(objectId, username, email string, roles []string) (string, error) {
    claims := CustomClaims{
        ObjectId: objectId,
        Username: username,
        Email:    email,
        Roles:    roles,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secretKey))
}
```

### Clean Architecture Dependency Injection

```go
// Repository Layer
commonUserRepository := repository.NewRepository[models.User](userCollection)
userRepository := repository.NewUserRepository(userCollection)

// Service Layer (depends on repositories)
userService := services.NewUserService(*commonUserRepository, userRepository)
authService := services.NewAuthService(userRepository, *commonUserRepository)

// Controller Layer (depends on services)
userController := controllers.NewUserController(userService)
authController := controllers.NewAuthController(authService)

// Routes (depends on controllers)
routes.UserRoutes(r, rateLimiter, *userController)
routes.AuthRoutes(r, rateLimiter, *authController)
```

---

## 🐳 Docker Configuration

### Multi-Stage Build

The `Dockerfile` uses a two-stage build process to minimize production image size:

```dockerfile
# Stage 1: Build (full Go toolchain)
FROM golang:1.22.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Stage 2: Runtime (minimal Alpine image)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

**Result**: 93% size reduction (400MB → 25MB)

### Docker Compose Services

```yaml
services:
  app:              # Go application (port 8080)
  mongo:            # Production MongoDB (port 27018)
  mongo-test:       # Test MongoDB (port 27019)
```

---

## 📄 Detailed Documentation

For comprehensive technical documentation including architecture diagrams, design decisions, and infrastructure details, see:

**[project-documentation.json](./project-documentation.json)** - Complete project specification with:
- Code showcase examples with explanations
- Infrastructure deployment layers
- Architecture patterns and decisions
- Data flow diagrams
- Technology decision records
- Scalability and security strategies

---

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Maintain clean architecture layer separation
- Add unit tests for new services
- Update Swagger documentation for new endpoints
- Keep DTOs and entities separate

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 📧 Contact & Support

- **Repository**: [github.com/aTrejoCoder/reading_tracker_api](https://github.com/aTrejoCoder/reading_tracker_api)
- **API Documentation**: http://localhost:8080/swagger/index.html
- **Issues**: [GitHub Issues](https://github.com/aTrejoCoder/reading_tracker_api/issues)

---

<p align="center">
  <strong>Built with ❤️ using Go, MongoDB, and Clean Architecture</strong>
</p>
