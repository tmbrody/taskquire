-- name: CreateTask :execresult
INSERT INTO tasks(id, name, description, created_at, updated_at, project_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetTaskByName :one
SELECT * FROM tasks
WHERE name = ?;

-- name: GetTasksByProjectID :many
SELECT * FROM tasks
WHERE project_id = ?;

-- name: UpdateTask :execresult
UPDATE tasks
SET name = ?, description = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteTask :execresult
DELETE FROM tasks
WHERE id = ?;