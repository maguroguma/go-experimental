// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING id, name
`

func (q *Queries) CreateUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, name)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, name FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}
