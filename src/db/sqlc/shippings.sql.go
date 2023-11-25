// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: shippings.sql

package db

import (
	"context"
)

const createShipping = `-- name: CreateShipping :one
INSERT INTO shippings (
  to_address,
  total
) VALUES (
$1, $2
)
RETURNING id, to_address, total
`

type CreateShippingParams struct {
	ToAddress string  `json:"to_address"`
	Total     float64 `json:"total"`
}

func (q *Queries) CreateShipping(ctx context.Context, arg CreateShippingParams) (Shipping, error) {
	row := q.db.QueryRowContext(ctx, createShipping, arg.ToAddress, arg.Total)
	var i Shipping
	err := row.Scan(&i.ID, &i.ToAddress, &i.Total)
	return i, err
}
