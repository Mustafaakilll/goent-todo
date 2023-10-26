package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/mustafaakilll/ent_todo/database"
	"github.com/mustafaakilll/ent_todo/handler"
	"github.com/mustafaakilll/ent_todo/middleware"
	"github.com/mustafaakilll/ent_todo/repository"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/mustafaakilll/ent_todo/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title           Ent Todo API
// @version         2.0
// @description     Basic implementation of swagger for Ent Todo API
// @termsOfService  http://swagger.io/terms/

// @license.name Apache 2.0

// @contact.name   Mustafa Akil
// @contact.email  mustafa@veriyaz.com.tr

// @host      localhost:3000
// @BasePath  /api/v1

// @accept json
// @produce json

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
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
	apiEndpoint := r.Group("api/v1")

	apiEndpoint.GET("/todos", middleware.Auth(), todoHandler.GetTodosHandler)
	apiEndpoint.GET("/todos/:id", middleware.Auth(), todoHandler.GetTodoByIdHandler)
	apiEndpoint.POST("/todos", middleware.Auth(), todoHandler.CreateTodosHandler)
	apiEndpoint.PUT("/todos/:id", middleware.Auth(), todoHandler.UpdateTodosHandler)
	apiEndpoint.DELETE("/todos/:id", middleware.Auth(), todoHandler.DeleteTodosHandler)

	apiEndpoint.POST("/register", userHandler.HandleRegisterUser)
	apiEndpoint.POST("/login", userHandler.HandleLoginUser)

	apiEndpoint.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	PORT := ":" + os.Getenv("PORT")

	r.Run(PORT)
}
