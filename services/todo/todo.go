package todo

import (
	"context"
)

type Todo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
}

type Repository interface {
	CreateTodo(ctx context.Context, todo Todo) error
	GetTodoByID(ctx context.Context, id string) (Todo, error)
	ChangeTodoStatus(ctx context.Context, id string, status int) error
}
