// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user_data (
    name,
    weight,
    height,
    age
) VALUES (
    $1, $2, $3, $4
) RETURNING id, name, weight, height, age, created_at
`
type CreateUserParams struct {
	Name string `json:"name"`
	Weight string `json:"weight"`
	Height string `json:"height"`
	Age string `json:"age"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (UserData, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Weight, arg.Height, arg.Age)
	var i UserData
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Weight,
		&i.Height,
		&i.Age,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM user_data
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, weight, height, age, created_at FROM user_data 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (UserData, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i UserData
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Weight,
		&i.Height,
		&i.Age,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, weight, height, age, created_at FROM user_data
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]UserData, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserData
	for rows.Next() {
		var i UserData
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Weight,
			&i.Height,
			&i.Age,
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

const updateUser = `-- name: UpdateUser :one
UPDATE user_data
SET weight = $2
WHERE id = $1
RETURNING id, name, weight, height, age, created_at
`

type UpdateUserParams struct {
	ID      int64 `json:"id"`
	Weight string `json:"weight"`
}


func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UserData, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.ID, arg.Weight)
	var i UserData
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Weight,
		&i.Height,
		&i.Age,
		&i.CreatedAt,
	)
	return i, err
}