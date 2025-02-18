-- name: CreateFeed :one
INSERT INTO feed (id, created_at, updated_at, name, url, user_id)
VALUES(
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feed;

-- name: GetFeedsWithUserName :many
SELECT
  feed.id,
  feed.name,
  feed.url,
  users.name as username
FROM feed
INNER JOIN users
ON users.id = feed.user_id;
