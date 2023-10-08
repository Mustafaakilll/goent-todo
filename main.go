package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mustafaakilll/ent_todo/database"
	"github.com/mustafaakilll/ent_todo/handler"
	"github.com/mustafaakilll/ent_todo/middleware"
	"github.com/mustafaakilll/ent_todo/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	client := database.Connect()

	repo := repository.NewTodoRepository(client)
	todoHandler := handler.NewTodoHandler(*repo)
	userHandler := handler.NewUserHandler(*repo)

	defer client.Close()

	r := gin.Default()

	r.GET("/todos", middleware.Auth(), todoHandler.GetTodosHandler)
	r.GET("/todos/:id", middleware.Auth(), todoHandler.GetTodoByIdHandler)
	r.POST("/todos", middleware.Auth(), todoHandler.CreateTodoHandler)
	r.PUT("/todos", middleware.Auth(), todoHandler.UpdateTodoByEntityHandler)
	r.DELETE("/todos/:id", middleware.Auth(), todoHandler.DeleteTodoHandler)

	r.POST("/register", userHandler.HandleRegisterUser)
	r.POST("/login", userHandler.HandleLoginUser)
	PORT := ":" + os.Getenv("PORT")

	r.Run(PORT)
}
