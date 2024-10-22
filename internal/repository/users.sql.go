// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username, email, password_hash
) VALUES (
    $1, $2, $3
)
RETURNING id, username, email, password_hash, created_at, updated_at, last_login, active_account
`

type CreateUserParams struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.ActiveAccount,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const disableUser = `-- name: DisableUser :exec
UPDATE users
SET active_account=FALSE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) DisableUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, disableUser, id)
	return err
}

const enableUser = `-- name: EnableUser :exec
UPDATE users
SET active_account=TRUE, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) EnableUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, enableUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at, last_login, active_account
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.ActiveAccount,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, created_at, updated_at, last_login, active_account
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.ActiveAccount,
	)
	return i, err
}

const getUserIdByEmail = `-- name: GetUserIdByEmail :one
SELECT id
FROM users
WHERE email = $1
`

func (q *Queries) GetUserIdByEmail(ctx context.Context, email string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getUserIdByEmail, email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserIdByUsername = `-- name: GetUserIdByUsername :one
SELECT id
FROM users
WHERE username = $1
`

func (q *Queries) GetUserIdByUsername(ctx context.Context, username string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getUserIdByUsername, username)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET username = $2, email = $3, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.Username, arg.Email)
	return err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users
SET email = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserEmailParams struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.Exec(ctx, updateUserEmail, arg.ID, arg.Email)
	return err
}

const updateUserLastLogin = `-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) UpdateUserLastLogin(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, updateUserLastLogin, id)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID           uuid.UUID `json:"id"`
	PasswordHash string    `json:"password_hash"`
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.ID, arg.PasswordHash)
	return err
}

const updateUserUsername = `-- name: UpdateUserUsername :exec
UPDATE users
SET username = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
`

type UpdateUserUsernameParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

func (q *Queries) UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) error {
	_, err := q.db.Exec(ctx, updateUserUsername, arg.ID, arg.Username)
	return err
}
