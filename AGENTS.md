# Agent Instructions for go-lab

## Project Context

This is a personal Go learning workspace. Purpose:
- Re-learn Go after not using it since 2020
- Experiment with different Go patterns and libraries
- Low-pressure, exploratory development

## Build & Run Commands

### getting-started
```bash
cd getting-started && go run main.go
```

### api
```bash
cd api && go run main.go
```

### Testing
```bash
cd <module> && go test ./...
```

### Vet & Format
```bash
cd <module> && go vet ./...
cd <module> && go fmt ./...
```

## Code Conventions

- Error handling: Return errors early, handle with `if err != nil`
- Context: Pass `context.Context` for cancellation and timeouts
- Modules: Each subdirectory is an independent Go module
- Packages: `package main` for executables

## Current Issues

### api/main.go:46 - Bug in todo creation

The current code has invalid syntax on line 46:
```go
err = gorm.G[Todos](db).Create(ctx, &todos)
```

Should be fixed to:
```go
if err := db.WithContext(ctx).Create(&todos).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
}
```

**Note:** The `err` variable shadowing issue also exists - the `err` from `DatabaseConnection` is used but the todo creation error is never checked properly.
