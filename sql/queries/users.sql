-- name: CreateUser :execresult
INSERT INTO users(id, name, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllUsers :many
SELECT * FROM users;