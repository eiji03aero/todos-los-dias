package todo

import (
	"context"
	"time"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository interface {
	GetTodos(ctx context.Context) ([]Todo, error)
	CreateTodo(ctx context.Context, todo Todo) error
	GetTodoByID(ctx context.Context, id string) (Todo, error)
	ChangeTodoStatus(ctx context.Context, id string, status int) error
}
