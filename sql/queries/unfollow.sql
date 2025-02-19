-- name: Unfollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feeds.id = feed_follows.feed_id
AND feed_follows.user_id = $1
AND feeds.url = $2;
