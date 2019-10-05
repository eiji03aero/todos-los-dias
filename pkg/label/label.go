package label

import (
	"context"
	"time"
)

type Label struct {
	ID        string    `json:"id"`
	Name      string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type LabelOwner struct {
	LabelId string `json:"label_id"`
	OwnerId string `json:"owner_id"`
}

type Repository interface {
	CreateLabel(ctx context.Context, label Label) error
}
