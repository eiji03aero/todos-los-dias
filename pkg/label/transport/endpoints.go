package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/eiji03aero/todos-los-dias/pkg/label"
)

type Endpoints struct {
	CreateLabel endpoint.Endpoint
}

func MakeEndpoints(s label.Service) Endpoints {
	return Endpoints{
		CreateLabel: makeCreateLabelEndpoint(s),
	}
}

func makeCreateLabelEndpoint(s label.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateLabelRequest)
		id, err := s.CreateLabel(ctx, req.Label)
		return CreateLabelResponse{ID: id, Err: err}, nil
	}
}
