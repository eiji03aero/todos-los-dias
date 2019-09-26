package todo

import (
	"context"
)

type Service interface {
	Index(ctx context.Context) (string, error)
}

type todoService struct{}

func NewService() Service {
	return &todoService{}
}

func (todoService) Index(ctx context.Context) (string, error) {
	return "todos will be here", nil
}
