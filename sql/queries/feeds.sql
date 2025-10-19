-- name: CreateFeed :one
INSERT INTO feeds(id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetFeed :one
SELECT * FROM feeds
WHERE id = $1 LIMIT 1;

-- name: GetAllFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds 
SET last_fetched_at = NOW(), updated_at = NOW()
WHERE id = $1
RETURNING *;