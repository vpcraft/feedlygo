-- name: CreateUser :one

INSERT INTO users(id, created_at, updated_at, fullname)
VALUES ($1, $2, $3, $4)
RETURNING *;