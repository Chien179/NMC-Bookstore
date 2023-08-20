// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: orders.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    username
) VALUES (
  $1
)
RETURNING id, username, created_at, status, sub_amount, sub_total
`

func (q *Queries) CreateOrder(ctx context.Context, username string) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, username)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.Status,
		&i.SubAmount,
		&i.SubTotal,
	)
	return i, err
}

const deleteOrder = `-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1
`

func (q *Queries) DeleteOrder(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOrder, id)
	return err
}

const getOrder = `-- name: GetOrder :one
SELECT id, username, created_at, status, sub_amount, sub_total FROM orders
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.Status,
		&i.SubAmount,
		&i.SubTotal,
	)
	return i, err
}

const getOrderToPayment = `-- name: GetOrderToPayment :one
SELECT id, username, created_at, status, sub_amount, sub_total FROM orders
WHERE username = $1 
LIMIT 1
`

func (q *Queries) GetOrderToPayment(ctx context.Context, username string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderToPayment, username)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.Status,
		&i.SubAmount,
		&i.SubTotal,
	)
	return i, err
}

const listOders = `-- name: ListOders :one
SELECT t.total_page, JSON_AGG(json_build_object
    ('id',t.id,
    'username',t.username,
    'status',t.status,
    'sub_amount',t.sub_amount,
    'sub_total',t.sub_total,
    'created_at',t.created_at)
    ) AS orders
	FROM (
      SELECT 
        CEILING(CAST(COUNT(id) OVER () AS FLOAT)/$1) AS total_page, id, username, created_at, status, sub_amount, sub_total 
      FROM orders
      ORDER BY id
      LIMIT $1
      OFFSET $2
    ) AS t
    GROUP BY t.total_page
`

type ListOdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListOdersRow struct {
	TotalPage float64         `json:"total_page"`
	Orders    json.RawMessage `json:"orders"`
}

func (q *Queries) ListOders(ctx context.Context, arg ListOdersParams) (ListOdersRow, error) {
	row := q.db.QueryRowContext(ctx, listOders, arg.Limit, arg.Offset)
	var i ListOdersRow
	err := row.Scan(&i.TotalPage, &i.Orders)
	return i, err
}

const listOdersByUserName = `-- name: ListOdersByUserName :many
SELECT id, username, created_at, status, sub_amount, sub_total FROM orders
WHERE username = $1
ORDER BY id
`

func (q *Queries) ListOdersByUserName(ctx context.Context, username string) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOdersByUserName, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.CreatedAt,
			&i.Status,
			&i.SubAmount,
			&i.SubTotal,
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

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
SET 
  status = COALESCE($1, status),
  sub_amount = COALESCE($2, sub_amount),
  sub_total = COALESCE($3, sub_total)
WHERE 
  id = $4
RETURNING id, username, created_at, status, sub_amount, sub_total
`

type UpdateOrderParams struct {
	Status    sql.NullString  `json:"status"`
	SubAmount sql.NullInt32   `json:"sub_amount"`
	SubTotal  sql.NullFloat64 `json:"sub_total"`
	ID        int64           `json:"id"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder,
		arg.Status,
		arg.SubAmount,
		arg.SubTotal,
		arg.ID,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.Status,
		&i.SubAmount,
		&i.SubTotal,
	)
	return i, err
}
