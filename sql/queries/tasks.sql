-- name: CreateTask :execresult
INSERT INTO tasks(id, name, description, project_id, team_id, owner_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

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