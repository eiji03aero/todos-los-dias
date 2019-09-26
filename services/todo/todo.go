package todo

import "context"

// Order represents an order
type Todo struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Status      int64  `json:"Status"`
	CreatedAt   string `json:"created_at"`
}

// Repository describes the persistence on order model
type Repository interface {
	// CreateOrder(ctx context.Context, order Order) error
	// GetOrderByID(ctx context.Context, id string) (Order, error)
	// ChangeOrderStatus(ctx context.Context, id string, status string) error
}
