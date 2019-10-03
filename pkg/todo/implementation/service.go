package implementation

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	todosvc "github.com/eiji03aero/todos-los-dias/pkg/todo"
)

type service struct {
	repository todosvc.Repository
	logger     log.Logger
}

func NewService(rep todosvc.Repository, logger log.Logger) todosvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) GetTodos(ctx context.Context) ([]todosvc.Todo, error) {
	logger := log.With(s.logger, "method", "GetTodos")
	todos, err := s.repository.GetTodos(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return todos, todosvc.ErrQueryRepository
	}
	return todos, nil
}

func (s *service) Create(ctx context.Context, todo todosvc.Todo) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	todo.ID = id
	todo.Status = 0
	todo.CreatedAt = time.Now()

	if err := s.repository.CreateTodo(ctx, todo); err != nil {
		level.Error(logger).Log("err", err)
		return "", todosvc.ErrCmdRepository
	}
	return id, nil
}

func (s *service) GetByID(ctx context.Context, id string) (todosvc.Todo, error) {
	logger := log.With(s.logger, "method", "GetByID")
	todo, err := s.repository.GetTodoByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		if err == sql.ErrNoRows {
			return todo, todosvc.ErrNotFound
		}
		return todo, todosvc.ErrQueryRepository
	}
	return todo, nil
}

func (s *service) ChangeStatus(ctx context.Context, id string, status int) error {
	logger := log.With(s.logger, "method", "ChangeStatus")
	if err := s.repository.ChangeTodoStatus(ctx, id, status); err != nil {
		level.Error(logger).Log("err", err)
		return todosvc.ErrCmdRepository
	}
	return nil
}
