// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: cities.sql

package db

import (
	"context"
)

const getCity = `-- name: GetCity :one
SELECT id, name, created_at FROM cities
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCity(ctx context.Context, id int64) (City, error) {
	row := q.db.QueryRowContext(ctx, getCity, id)
	var i City
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listCities = `-- name: ListCities :many
SELECT id, name, created_at FROM cities
ORDER BY id
`

func (q *Queries) ListCities(ctx context.Context) ([]City, error) {
	rows, err := q.db.QueryContext(ctx, listCities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []City{}
	for rows.Next() {
		var i City
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
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
