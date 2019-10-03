package cockroachdb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log/level"

	"github.com/cockroachdb/cockroach-go/crdb"
	"github.com/go-kit/kit/log"

	"github.com/eiji03aero/todos-los-dias/pkg/todo"
)

var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *sql.DB, logger log.Logger) (todo.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "cockroachdb"),
	}, nil
}

func (repo *repository) GetTodos(ctx context.Context) ([]todo.Todo, error) {
	var todos []todo.Todo
	stmt, err := repo.db.Prepare("SELECT id, title, description, status, created_at FROM todos")
	if err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return todos, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return todos, err
	}
	for rows.Next() {
		var t todo.Todo
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		todos = append(todos, t)
	}
	if err := rows.Err(); err != nil {
		level.Error(repo.logger).Log("err", err)
		return todos, err
	}
	return todos, nil
}

func (repo *repository) CreateTodo(ctx context.Context, todo todo.Todo) error {
	// Run a transaction to sync the query model.
	err := crdb.ExecuteTx(ctx, repo.db, nil, func(tx *sql.Tx) error {
		return createTodo(tx, todo)
	})
	if err != nil {
		return err
	}
	return nil
}

func createTodo(tx *sql.Tx, todo todo.Todo) error {
	sql := `
		INSERT INTO todos (id, title, description, status, created_at)
		VALUES ($1,$2,$3,$4,$5)`
	_, err := tx.Exec(sql, todo.ID, todo.Title, todo.Description, todo.Status, todo.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) ChangeTodoStatus(ctx context.Context, todoId string, status int) error {
	sql := `
		UPDATE todos
		SET status=$2
		WHERE id=$1`

	_, err := repo.db.ExecContext(ctx, sql, todoId, status)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetTodoByID(ctx context.Context, id string) (todo.Todo, error) {
	var todoRow = todo.Todo{}
	if err := repo.db.QueryRowContext(ctx,
		"SELECT id, title, description, status FROM todos WHERE id = $1",
		id).
		Scan(
			&todoRow.ID, &todoRow.Title, &todoRow.Description, &todoRow.Status,
		); err != nil {
		level.Error(repo.logger).Log("err", err.Error())
		return todoRow, err
	}
	return todoRow, nil
}

func (repo *repository) Close() error {
	return repo.db.Close()
}
