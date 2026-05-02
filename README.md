# go-lab

A personal Go learning and experimentation workspace. This project serves as a **refresh/reboot journey** to re-learn Go programming after not using it since 2020.

## Project Structure

```
go-lab/
├── getting-started/    # Simplest possible Go module (Hello World)
└── api/                # REST API with Gin + GORM + SQLite
```

## Modules

### getting-started

Minimal Go module demonstrating basic syntax. Prints "Hello world!" to stdout.

```bash
cd getting-started
go run main.go
```

### api

REST API demonstrating:
- **Gin** - Web framework
- **GORM** - ORM for database operations
- **SQLite** - Lightweight database

**Endpoints:**
- `GET /ping` - Health check, returns `{"message": "pong"}`
- `POST /todos` - Create a todo item

```bash
cd api
go run main.go
# Server runs on http://localhost:8080
```

## Topics to Explore

This project covers a refresh of:

- Go modules and workspace structure
- REST API development with Gin
- Database integration with GORM and SQLite
- Context management (`context.Context`)
- JSON request/response handling
- Basic CRUD operations

## Requirements

- Go 1.21+
- (api module) SQLite3
