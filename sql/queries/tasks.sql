-- name: CreateTask :execresult
INSERT INTO tasks(id, name, description, created_at, updated_at, project_id)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllTasks :many
SELECT * FROM tasks;