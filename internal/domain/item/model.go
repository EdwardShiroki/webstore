package item

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	CreatedAt   string    `json:"created_at"`
}
