package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/repository"
)

type TodoHandler struct {
	TodoRepository repository.TodoRepository
}

func NewTodoHandler(repository repository.TodoRepository) *TodoHandler {
	return &TodoHandler{
		TodoRepository: repository,
	}
}

func (uh TodoHandler) GetTodosHandler(ctx *gin.Context) {

	todos, err := uh.TodoRepository.GetTodos()
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (uh TodoHandler) GetTodoByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)

	todos, err := uh.TodoRepository.GetTodoById(todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func (uh TodoHandler) CreateTodosHandler(ctx *gin.Context) {

	var todo ent.Todo
	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.MustGet("user_id")

	todos, err := uh.TodoRepository.CreateTodo(todo, userId.(uuid.UUID))
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}

func (uh TodoHandler) DeleteTodosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)

	err = uh.TodoRepository.DeleteTodo(todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": "Todo deleted successfully",
	})
}

func (uh TodoHandler) UpdateTodosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)
	var todo ent.Todo

	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo, err := uh.TodoRepository.UpdateTodo(todo, todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newTodo,
	})
}
