-- name: CreateTask :execresult
INSERT INTO tasks(id, name, description, created_at, updated_at, project_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllTasks :many
SELECT * FROM tasks;

-- name: GetTaskById :one
SELECT * FROM tasks
WHERE id = ?;

-- name: GetTasksByProjectID :many
SELECT * FROM tasks
WHERE project_id = ?;