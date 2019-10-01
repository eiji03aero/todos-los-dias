package todo

import (
	"context"
	"errors"
)

var (
	ErrNotFound        = errors.New("todo not found")
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	GetTodos(ctx context.Context) ([]Todo, error)
	Create(ctx context.Context, todo Todo) (string, error)
	GetByID(ctx context.Context, id string) (Todo, error)
	ChangeStatus(ctx context.Context, id string, status int) error
}
