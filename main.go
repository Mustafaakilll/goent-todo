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

	todoRepo := repository.NewTodoRepository(client)
	userRepo := repository.NewUserRepository(client)
	todoHandler := handler.NewTodoHandler(*todoRepo)
	userHandler := handler.NewUserHandler(*userRepo)

	defer client.Close()

	r := gin.Default()

	r.GET("/todos", middleware.Auth(), todoHandler.GetTodosHandler)
	r.GET("/todos/:id", middleware.Auth(), todoHandler.GetTodoByIdHandler)
	r.POST("/todos", middleware.Auth(), todoHandler.CreateTodosHandler)
	r.PUT("/todos/:id", middleware.Auth(), todoHandler.UpdateTodosHandler)
	r.DELETE("/todos/:id", middleware.Auth(), todoHandler.DeleteTodosHandler)
	//
	r.POST("/register", userHandler.HandleRegisterUser)
	r.POST("/login", userHandler.HandleLoginUser)
	PORT := ":" + os.Getenv("PORT")

	r.Run(PORT)
}
