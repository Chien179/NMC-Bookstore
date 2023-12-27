// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: dislike.sql

package db

import (
	"context"
)

const createdDislike = `-- name: CreatedDislike :one
INSERT INTO "dislike" (username, review_id, is_dislike)
VALUES ($1, $2, $3)
RETURNING id, username, review_id, is_dislike
`

type CreatedDislikeParams struct {
	Username  string `json:"username"`
	ReviewID  int64  `json:"review_id"`
	IsDislike bool   `json:"is_dislike"`
}

func (q *Queries) CreatedDislike(ctx context.Context, arg CreatedDislikeParams) (Dislike, error) {
	row := q.db.QueryRowContext(ctx, createdDislike, arg.Username, arg.ReviewID, arg.IsDislike)
	var i Dislike
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ReviewID,
		&i.IsDislike,
	)
	return i, err
}

const getDislike = `-- name: GetDislike :one
SELECT id, username, review_id, is_dislike
FROM "dislike"
WHERE username = $1
    AND review_id = $2
LIMIT 1
`

type GetDislikeParams struct {
	Username string `json:"username"`
	ReviewID int64  `json:"review_id"`
}

func (q *Queries) GetDislike(ctx context.Context, arg GetDislikeParams) (Dislike, error) {
	row := q.db.QueryRowContext(ctx, getDislike, arg.Username, arg.ReviewID)
	var i Dislike
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ReviewID,
		&i.IsDislike,
	)
	return i, err
}

const listDislike = `-- name: ListDislike :many
SELECT id, username, review_id, is_dislike
FROM "dislike"
WHERE username = $1
ORDER BY review_id
`

func (q *Queries) ListDislike(ctx context.Context, username string) ([]Dislike, error) {
	rows, err := q.db.QueryContext(ctx, listDislike, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Dislike{}
	for rows.Next() {
		var i Dislike
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.ReviewID,
			&i.IsDislike,
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

const updateDislike = `-- name: UpdateDislike :one
UPDATE "dislike"
SET is_dislike = $1
WHERE username = $2
    AND review_id = $3
RETURNING id, username, review_id, is_dislike
`

type UpdateDislikeParams struct {
	IsDislike bool   `json:"is_dislike"`
	Username  string `json:"username"`
	ReviewID  int64  `json:"review_id"`
}

func (q *Queries) UpdateDislike(ctx context.Context, arg UpdateDislikeParams) (Dislike, error) {
	row := q.db.QueryRowContext(ctx, updateDislike, arg.IsDislike, arg.Username, arg.ReviewID)
	var i Dislike
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.ReviewID,
		&i.IsDislike,
	)
	return i, err
}
