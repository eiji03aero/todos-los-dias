package implementation

import (
	"context"
	_ "database/sql"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	labelsvc "github.com/eiji03aero/todos-los-dias/pkg/label"
)

type service struct {
	repository labelsvc.Repository
	logger     log.Logger
}

func NewService(rep labelsvc.Repository, logger log.Logger) labelsvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) CreateLabel(ctx context.Context, label labelsvc.Label) (string, error) {
	logger := log.With(s.logger, "method", "CreateLabel")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	label.ID = id
	label.CreatedAt = time.Now()

	if err := s.repository.CreateLabel(ctx, label); err != nil {
		level.Error(logger).Log("err", err)
		return "", labelsvc.ErrCmdRepository
	}

	return id, nil
}
