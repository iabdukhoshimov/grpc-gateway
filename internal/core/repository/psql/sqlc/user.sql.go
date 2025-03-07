// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package sqlc

import (
	"context"
)

const countUser = `-- name: CountUser :one
SELECT count(*) FROM "user"
`

func (q *Queries) CountUser(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :exec
INSERT INTO "user" (name, email) VALUES ($1, $2)
`

type CreateUserParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser, arg.Name, arg.Email)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT id, name, email, created_at, updated_at, deleted_at FROM "user"
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, name, email, created_at, updated_at, deleted_at FROM "user"
WHERE name = $1
`

func (q *Queries) GetUserByName(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByName, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, name, email, created_at, updated_at, deleted_at FROM "user"
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :exec
UPDATE "user"
SET name = $2, email = $3
WHERE id = $1
`

type UpdateUserParams struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.Name, arg.Email)
	return err
}
