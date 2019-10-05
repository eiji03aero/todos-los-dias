package cockroachdb

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/go-kit/kit/log/level"

	_ "github.com/cockroachdb/cockroach-go/crdb"
	"github.com/go-kit/kit/log"

	"github.com/eiji03aero/todos-los-dias/pkg/label"
)

var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func New(db *sql.DB, logger log.Logger) (label.Repository, error) {
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "cockroachdb"),
	}, nil
}

func (repo *repository) CreateLabel(ctx context.Context, label label.Label) error {
	sql := `
		INSERT INTO labels (id, name, created_at)
		VALUES ($1, $2, $3)
	`
	_, err := repo.db.ExecContext(ctx, sql, label.ID, label.Name, label.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
