package transport

import (
	"github.com/eiji03aero/todos-los-dias/pkg/label"
)

type CreateLabelRequest struct {
	Label label.Label
}

type CreateLabelResponse struct {
	ID  string
	Err error
}
