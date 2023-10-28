-- name: CreateUser :execresult
INSERT INTO users(id, name, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :execresult
UPDATE users
SET name = ?, email = ?, password = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteUser :execresult
DELETE FROM users
WHERE id = ?;