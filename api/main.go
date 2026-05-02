package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := DatabaseConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		// return JSON response
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/todos", func(c *gin.Context) {
		var todos Todos
		if err := c.ShouldBindJSON(&todos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Println("title", todos.Title)
		log.Println("description", todos.Description)

		// insert into database
		err = gorm.G[Todos](db).Create(ctx, &todos)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func DatabaseConnection(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema ...
	err = db.WithContext(ctx).AutoMigrate(&Todos{})

	return db, err
}
