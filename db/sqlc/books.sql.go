// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: books.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (
  name,
  price,
  image,
  description,
  author,
  publisher,
  quantity
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, name, price, image, description, author, publisher, quantity, created_at
`

type CreateBookParams struct {
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Image       []string `json:"image"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Publisher   string   `json:"publisher"`
	Quantity    int32    `json:"quantity"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Name,
		arg.Price,
		pq.Array(arg.Image),
		arg.Description,
		arg.Author,
		arg.Publisher,
		arg.Quantity,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		pq.Array(&i.Image),
		&i.Description,
		&i.Author,
		&i.Publisher,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT id, name, price, image, description, author, publisher, quantity, created_at FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		pq.Array(&i.Image),
		&i.Description,
		&i.Author,
		&i.Publisher,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT id, name, price, image, description, author, publisher, quantity, created_at FROM books
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Price,
			pq.Array(&i.Image),
			&i.Description,
			&i.Author,
			&i.Publisher,
			&i.Quantity,
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

const updateBook = `-- name: UpdateBook :one
UPDATE books
SET name = COALESCE($1, name),
  price = COALESCE($2, price),
  image = COALESCE($3, image),
  description = COALESCE($4, description),
  author = COALESCE($5, author),
  publisher = COALESCE($6, publisher),
  quantity = COALESCE($7, quantity)
WHERE id = $8
RETURNING id, name, price, image, description, author, publisher, quantity, created_at
`

type UpdateBookParams struct {
	Name        sql.NullString  `json:"name"`
	Price       sql.NullFloat64 `json:"price"`
	Image       []string        `json:"image"`
	Description sql.NullString  `json:"description"`
	Author      sql.NullString  `json:"author"`
	Publisher   sql.NullString  `json:"publisher"`
	Quantity    sql.NullInt32   `json:"quantity"`
	ID          int64           `json:"id"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook,
		arg.Name,
		arg.Price,
		pq.Array(arg.Image),
		arg.Description,
		arg.Author,
		arg.Publisher,
		arg.Quantity,
		arg.ID,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Price,
		pq.Array(&i.Image),
		&i.Description,
		&i.Author,
		&i.Publisher,
		&i.Quantity,
		&i.CreatedAt,
	)
	return i, err
}
