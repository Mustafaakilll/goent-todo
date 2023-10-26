package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mustafaakilll/ent_todo/ent"
	"github.com/mustafaakilll/ent_todo/ent/todo"
)

// TodoRepository struct for accessing Ent.Client
type TodoRepository struct {
	// ent Client
	Client *ent.Client
}

// NewTodoRepository function for creating new TodoRepository with client for DI
func NewTodoRepository(client *ent.Client) *TodoRepository {
	return &TodoRepository{
		Client: client,
	}
}

// GetTodos function for getting all todos of user
func (r TodoRepository) GetTodos() ([]*ent.Todo, error) {
	ctx := context.Background()
	todos, err := r.Client.Todo.Query().All(ctx)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetTodoById function for getting todo by id of user.
//
// You have to provide todoId
func (r TodoRepository) GetTodoById(todoId uuid.UUID) (*ent.Todo, error) {
	ctx := context.Background()
	todo, err := r.Client.Todo.Query().Where(todo.ID(todoId)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// CreateTodo function for creating new todo for user.
//
// You have to provide newTodo and userId.
//
// NewTodo is ent.Todo struct
func (r TodoRepository) CreateTodo(newTodo ent.Todo, userId uuid.UUID) (*ent.Todo, error) {
	ctx := context.Background()
	todo, err := r.Client.Todo.Create().
		SetTitle(newTodo.Title).
		SetUserID(userId).
		SetDescription(newTodo.Description).
		SetNillableDueDate(newTodo.DueDate).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo function for deleting todo by id of user.
//
// You have to provide todoId.
func (r TodoRepository) DeleteTodo(todoId uuid.UUID) error {
	fmt.Println(todoId)
	ctx := context.Background()
	err := r.Client.Todo.DeleteOneID(todoId).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

// UpdateTodo function for updating todo by id of user.
//
// You have to provide todoId and which fields have to update.
func (r TodoRepository) UpdateTodo(newTodo ent.Todo, todoId uuid.UUID) (*ent.Todo, error) {
	ctx := context.Background()
	oldTodo, err := r.GetTodoById(todoId)
	_newTodo := ValidateTodo(oldTodo, &newTodo)

	err = r.Client.Todo.
		UpdateOneID(todoId).
		SetTitle(_newTodo.Title).
		SetDescription(_newTodo.Description).
		SetNillableDueDate(_newTodo.DueDate).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return _newTodo, nil
}

// ValidateTodo function for validating todo fields.
//
// Checking which fields have to update.
// If newTodo fields are empty, oldTodo fields will be used.
func ValidateTodo(oldTodo, newTodo *ent.Todo) *ent.Todo {
	if newTodo.Title != "" {
		oldTodo.Title = newTodo.Title
	}

	if newTodo.Description != "" {
		oldTodo.Description = newTodo.Description
	}

	if newTodo.DueDate != nil {
		oldTodo.DueDate = newTodo.DueDate
	}

	return oldTodo
}
