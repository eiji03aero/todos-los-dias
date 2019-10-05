package label

import (
	"context"
	"errors"
)

var (
	ErrCmdRepository   = errors.New("unable to command repository")
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	CreateLabel(ctx context.Context, label Label) (string, error)
}
