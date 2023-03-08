// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: carts.sql

package db

import (
	"context"
)

const createCart = `-- name: CreateCart :one
INSERT INTO carts (
  books_id,
  username
) VALUES (
  $1, $2
)
RETURNING id, books_id, username, created_at
`

type CreateCartParams struct {
	BooksID  int64  `json:"books_id"`
	Username string `json:"username"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, createCart, arg.BooksID, arg.Username)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.BooksID,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE FROM carts
WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}

const getCart = `-- name: GetCart :one
SELECT id, books_id, username, created_at FROM carts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCart(ctx context.Context, id int64) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getCart, id)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.BooksID,
		&i.Username,
		&i.CreatedAt,
	)
	return i, err
}

const listCartsByUsername = `-- name: ListCartsByUsername :many
SELECT id, books_id, username, created_at FROM carts
WHERE username = $1
ORDER BY id
`

func (q *Queries) ListCartsByUsername(ctx context.Context, username string) ([]Cart, error) {
	rows, err := q.db.QueryContext(ctx, listCartsByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Cart{}
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.ID,
			&i.BooksID,
			&i.Username,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
