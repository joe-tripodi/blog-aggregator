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

-- name: GetFeedFollowsForUser :many
SELECT
  feed_follows.*,
  users.name AS user_name,
  feeds.name AS feed_name
FROM feed_follows
INNER JOIN users
ON users.id = feed_follows.user_id
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
WHERE $1 = feed_follows.user_id;

-- name: CreateFeedFollows :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES ($1, $2, $3, $4, $5) RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds
ON feeds.id = inserted_feed_follow.feed_id;

