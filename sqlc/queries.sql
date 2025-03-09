-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING id, name;

-- name: GetUser :one
SELECT id, name FROM users WHERE id = $1;
