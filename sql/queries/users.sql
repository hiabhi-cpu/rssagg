-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, user_name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE user_name LIKE $1 LIMIT 1;

-- name: DeleteAll :exec
DELETE FROM users;

-- name: ListUsers :many
SELECT user_name from users
ORDER BY user_name;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1 LIMIT 1;