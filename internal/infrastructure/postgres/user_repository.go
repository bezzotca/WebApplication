// Package postgres (infrastructure) implements the usecase repository ports
// on top of the connection pool produced by the isolated pkg/database/postgres
// core. This is the only place SQL and pgx types are allowed to appear.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"webproject/internal/domain"
	"webproject/internal/domain/entity"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, u *entity.User) error {
	const q = `INSERT INTO users (email, name) VALUES ($1, $2) RETURNING id, created_at`
	return r.pool.QueryRow(ctx, q, u.Email, u.Name).Scan(&u.ID, &u.CreatedAt)
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	const q = `SELECT id, email, name, created_at FROM users WHERE id = $1`

	var u entity.User
	err := r.pool.QueryRow(ctx, q, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}
