package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func openDB(ctx context.Context) (*pgxpool.Pool, error) {
	dsn := env("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/webproject?sslmode=disable")

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return pool, nil
}

func createUser(ctx context.Context, pool *pgxpool.Pool, u *User) error {
	const q = `INSERT INTO users (email, name) VALUES ($1, $2) RETURNING id, created_at`
	return pool.QueryRow(ctx, q, u.Email, u.Name).Scan(&u.ID, &u.CreatedAt)
}

func getUser(ctx context.Context, pool *pgxpool.Pool, id int64) (*User, error) {
	const q = `SELECT id, email, name, created_at FROM users WHERE id = $1`

	var u User
	if err := pool.QueryRow(ctx, q, id).Scan(&u.ID, &u.Email, &u.Name, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
