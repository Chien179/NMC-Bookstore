// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: reviews.sql

package db

import (
	"context"
	"encoding/json"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (
    username,
    books_id,
    comments,
    rating
  )
VALUES ($1, $2, $3, $4)
RETURNING id, username, books_id, comments, rating, created_at
`

type CreateReviewParams struct {
	Username string `json:"username"`
	BooksID  int64  `json:"books_id"`
	Comments string `json:"comments"`
	Rating   int32  `json:"rating"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRowContext(ctx, createReview,
		arg.Username,
		arg.BooksID,
		arg.Comments,
		arg.Rating,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.BooksID,
		&i.Comments,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = $1
`

func (q *Queries) DeleteReview(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReview, id)
	return err
}

const getReview = `-- name: GetReview :one
SELECT id, username, books_id, comments, rating, created_at
FROM reviews
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetReview(ctx context.Context, id int64) (Review, error) {
	row := q.db.QueryRowContext(ctx, getReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.BooksID,
		&i.Comments,
		&i.Rating,
		&i.CreatedAt,
	)
	return i, err
}

const listReviewsByBookID = `-- name: ListReviewsByBookID :one
SELECT t.total_page, JSON_AGG( json_build_object (
      'id', id,
      'username', username,
      'image', image,
      'books_id', books_id,
      'comments', comments,
      'rating', rating,
      'created_at', created_at
    )
  ) AS reviews
FROM (
    SELECT reviews.id,
      CEILING(CAST(COUNT(id) OVER () AS FLOAT) / $2) AS total_page,
      users.username AS username,
      users.image AS image,
      reviews.books_id AS books_id,
      reviews."comments" AS "comments",
      reviews.rating AS rating,
      reviews.created_at AS created_at
    FROM reviews
      INNER JOIN users ON reviews.username = users.username
    WHERE reviews.books_id = $1
    ORDER BY id
    LIMIT $2 OFFSET $3
  ) AS t
GROUP BY t.total_page
`

type ListReviewsByBookIDParams struct {
	BooksID int64 `json:"books_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

type ListReviewsByBookIDRow struct {
	TotalPage float64         `json:"total_page"`
	Reviews   json.RawMessage `json:"reviews"`
}

func (q *Queries) ListReviewsByBookID(ctx context.Context, arg ListReviewsByBookIDParams) (ListReviewsByBookIDRow, error) {
	row := q.db.QueryRowContext(ctx, listReviewsByBookID, arg.BooksID, arg.Limit, arg.Offset)
	var i ListReviewsByBookIDRow
	err := row.Scan(&i.TotalPage, &i.Reviews)
	return i, err
}
