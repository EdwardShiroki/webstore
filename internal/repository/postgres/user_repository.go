package postgres

import (
	"context"
	"database/sql"

	"github.com/EdwardShiroki/webstore/internal/domain/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) error {
	query := `
		INSERT INTO users (login, password_hash, role, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Login,
		user.PasswordHash,
		user.Role,
		user.CreatedAt,
	).Scan(&user.ID)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	query := `SELECT id, login, password_hash, role, created_at FROM users WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var u user.User
	err := row.Scan(&u.ID, &u.Login, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*user.User, error) {
	query := `SELECT id, login, password_hash, role, created_at FROM users WHERE login = $1`
	row := r.db.QueryRowContext(ctx, query, login)
	var u user.User
	err := row.Scan(&u.ID, &u.Login, &u.PasswordHash, &u.Role, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return &u, nil
}
