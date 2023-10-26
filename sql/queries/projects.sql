-- name: CreateProject :execresult
INSERT INTO projects(id, name, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: GetAllProjects :many
SELECT * FROM projects;