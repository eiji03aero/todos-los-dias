package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/eiji03aero/todos-los-dias/pkg/todo"
)

type Endpoints struct {
	GetTodos     endpoint.Endpoint
	Create       endpoint.Endpoint
	GetByID      endpoint.Endpoint
	ChangeStatus endpoint.Endpoint
}

func MakeEndpoints(s todo.Service) Endpoints {
	return Endpoints{
		GetTodos:     makeGetTodosEndpoint(s),
		Create:       makeCreateEndpoint(s),
		GetByID:      makeGetByIDEndpoint(s),
		ChangeStatus: makeChangeStatusEndpoint(s),
	}
}

func makeGetTodosEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		todos, err := s.GetTodos(ctx)
		return GetTodosResponse{Todos: todos, Err: err}, nil
	}
}

func makeCreateEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Todo)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetByIDEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		todoRes, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{Todo: todoRes, Err: err}, nil
	}
}

func makeChangeStatusEndpoint(s todo.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeStatusRequest)
		err := s.ChangeStatus(ctx, req.ID, req.Status)
		return ChangeStatusResponse{Err: err}, nil
	}
}
