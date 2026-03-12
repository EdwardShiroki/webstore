package app

import (
	"database/sql"
	"os"

	"github.com/EdwardShiroki/webstore/internal/repository/postgres"
	"github.com/EdwardShiroki/webstore/internal/service/auth"
	"github.com/EdwardShiroki/webstore/internal/transport/http/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Run() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.Default()

	userRepo := postgres.NewUserRepository(db)
	itemRepo := postgres.NewItemRepository(db)
	authService := auth.NewService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	r.POST("/register", authHandler.Register)
	r.GET("/items/:id", handler.NewItemHandler(itemRepo).GetByID)
	r.GET("/items", handler.NewItemHandler(itemRepo).List)
	r.POST("/items", handler.NewItemHandler(itemRepo).Create)

	r.Run(":8080")
}
