-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
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
SELECT * FROM feeds;

-- name: GetFeedsWithUserName :many
SELECT
  feeds.id,
  feeds.name,
  feeds.url,
  users.name as username
FROM feeds
INNER JOIN users
ON users.id = feeds.user_id;

-- name: GetFeedByUrl :one
SELECT * FROM feeds where $1 = feeds.url;

