-- name: CreateFeed :one
INSERT INTO feeds(id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;