CREATE TABLE posts (
    id      uuid    PRIMARY KEY,
    title   text    NOT NULL,
    body    text
);

-- name: GetAllPosts :many
SELECT * FROM posts
ORDER BY title;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: CreatePost :one
INSERT INTO posts (
    title, body
) VALUES (
    $1, $2
)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;