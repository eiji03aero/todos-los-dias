package transport

import (
	"github.com/eiji03aero/todos-los-dias/services/todo"
)

type GetTodosRequest struct {
}

type GetTodosResponse struct {
	Todos []todo.Todo `json:"todos"`
	Err   error       `json:"error,omitempty"`
}

type CreateRequest struct {
	Todo todo.Todo
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

type GetByIDRequest struct {
	ID string
}

type GetByIDResponse struct {
	Todo todo.Todo `json:"todo"`
	Err  error     `json:"error,omitempty"`
}

type ChangeStatusRequest struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type ChangeStatusResponse struct {
	Err error `json:"error,omitempty"`
}
