package item

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, item *Item) error
	GetByID(ctx context.Context, id uuid.UUID) (*Item, error)
	List(ctx context.Context) ([]*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, id uuid.UUID) error
}
