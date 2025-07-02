-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE user_id = $1 AND feed_id = $2;

-- name: GetFeedFollowsByUser :many
SELECT ff.*, f.name as feed_name, f.url as feed_url
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
WHERE ff.user_id = $1;
