// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  full_name,
  email,
  password,
  age,
  sex,
  image,
  phone_number
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING username, full_name, email, password, image, phone_number, age, sex, role, password_changed_at, created_at, is_email_verified
`

type CreateUserParams struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Age         int32  `json:"age"`
	Sex         string `json:"sex"`
	Image       string `json:"image"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.Password,
		arg.Age,
		arg.Sex,
		arg.Image,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Image,
		&i.PhoneNumber,
		&i.Age,
		&i.Sex,
		&i.Role,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, full_name, email, password, image, phone_number, age, sex, role, password_changed_at, created_at, is_email_verified FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Image,
		&i.PhoneNumber,
		&i.Age,
		&i.Sex,
		&i.Role,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT username, full_name, email, password, image, phone_number, age, sex, role, password_changed_at, created_at, is_email_verified FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Image,
		&i.PhoneNumber,
		&i.Age,
		&i.Sex,
		&i.Role,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT username, full_name, email, password, image, phone_number, age, sex, role, password_changed_at, created_at, is_email_verified FROM users
ORDER BY username
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.Email,
			&i.Password,
			&i.Image,
			&i.PhoneNumber,
			&i.Age,
			&i.Sex,
			&i.Role,
			&i.PasswordChangedAt,
			&i.CreatedAt,
			&i.IsEmailVerified,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET full_name = COALESCE($1, full_name),
    email = COALESCE($2, email),
    image = COALESCE($3, image),
    phone_number = COALESCE($4, phone_number),
    age = COALESCE($5, age),
    sex = COALESCE($6, sex),
    password = COALESCE($7, password),
    password_changed_at = COALESCE($8, password_changed_at),
    is_email_verified = COALESCE($9, is_email_verified)
WHERE 
  username = $10
RETURNING username, full_name, email, password, image, phone_number, age, sex, role, password_changed_at, created_at, is_email_verified
`

type UpdateUserParams struct {
	FullName          sql.NullString `json:"full_name"`
	Email             sql.NullString `json:"email"`
	Image             sql.NullString `json:"image"`
	PhoneNumber       sql.NullString `json:"phone_number"`
	Age               sql.NullInt32  `json:"age"`
	Sex               sql.NullString `json:"sex"`
	Password          sql.NullString `json:"password"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	IsEmailVerified   sql.NullBool   `json:"is_email_verified"`
	Username          string         `json:"username"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FullName,
		arg.Email,
		arg.Image,
		arg.PhoneNumber,
		arg.Age,
		arg.Sex,
		arg.Password,
		arg.PasswordChangedAt,
		arg.IsEmailVerified,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.Password,
		&i.Image,
		&i.PhoneNumber,
		&i.Age,
		&i.Sex,
		&i.Role,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}
