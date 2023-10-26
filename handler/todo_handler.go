package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/model"
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

// GetTodosHandler godoc
// @Summary      Get all todos
// @Description  Get all todos by user
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  ent.Todo
// @Failure      400  {object}  model.Response
// @Router       /todos [get]
func (uh TodoHandler) GetTodosHandler(ctx *gin.Context) {

	todos, err := uh.TodoRepository.GetTodos()
	if err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// GetTodoByIdHandler godoc
// @Summary      Get todo by id
// @Description  Get todo by id of user
// @Produce      json
// @Security     ApiKeyAuth
// @Param id path string true "TodoId"
// @Success      200  {object}  ent.Todo
// @Failure      400  {object}  model.Response
// @Router       /todos/{id} [get]
func (uh TodoHandler) GetTodoByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)

	todos, err := uh.TodoRepository.GetTodoById(todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

type CreateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// CreateTodosHandler godoc
// @Summary      Create Todo
// @Description  Create Todo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param todo body CreateTodo true "Todo"
// @Success      200  {object}  ent.Todo
// @Failure      400  {object}  model.Response
// @Router       /todos [post]
func (uh TodoHandler) CreateTodosHandler(ctx *gin.Context) {
	var todo ent.Todo
	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
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

// CreateTodosHandler godoc
// @Summary      Create Todo
// @Description  Create Todo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param todoId path string true "TodoID"
// @Success      200  {object} model.Response
// @Failure      400  {object} model.Response
// @Router       /todos/{id} [delete]
func (uh TodoHandler) DeleteTodosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)

	err = uh.TodoRepository.DeleteTodo(todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}

	ctx.JSON(model.NewResponse(http.StatusOK, "Todo deleted successfully"))
}

// UpdateTodosHandler godoc
// @Summary      Update Todo
// @Description  Upate Todo
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param todoId path string true "TodoID"
// @Param todo body CreateTodo true "Todo"
// @Success      200  {object} ent.Todo
// @Failure      400  {object} model.Response
// @Router       /todos/{id} [put]
func (uh TodoHandler) UpdateTodosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	todoId, err := uuid.Parse(id)
	var todo ent.Todo

	if err := ctx.ShouldBind(&todo); err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}

	newTodo, err := uh.TodoRepository.UpdateTodo(todo, todoId)
	if err != nil {
		ctx.Abort()
		ctx.JSON(model.NewResponse(http.StatusBadRequest, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": newTodo,
	})
}
