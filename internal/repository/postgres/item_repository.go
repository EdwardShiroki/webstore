package postgres

import (
	"context"
	"database/sql"

	"github.com/EdwardShiroki/webstore/internal/domain/item"
	"github.com/google/uuid"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(ctx context.Context, item *item.Item) error {
	query := `
		INSERT INTO items (name, description, price, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRowContext(
		ctx,
		query,
		item.Name,
		item.Description,
		item.Price,
		item.CreatedAt,
	).Scan(&item.ID)
	return err
}

func (r *ItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*item.Item, error) {
	query := `SELECT id, name, description, price, created_at FROM items WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var i item.Item
	err := row.Scan(&i.ID, &i.Name, &i.Description, &i.Price, &i.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (r *ItemRepository) List(ctx context.Context) ([]*item.Item, error) {
	query := `SELECT id, name, description, price, created_at FROM items`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []*item.Item
	for rows.Next() {
		var i item.Item
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Price, &i.CreatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, &i)
	}

	return items, nil
}

func (r *ItemRepository) Update(ctx context.Context, item *item.Item) error {
	query := `
		UPDATE items
		SET name = $1, description = $2, price = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		item.Name,
		item.Description,
		item.Price,
		item.ID,
	)
	return err
}

func (r *ItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM items WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
