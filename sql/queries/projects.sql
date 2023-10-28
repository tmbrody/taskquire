-- name: CreateProject :execresult
INSERT INTO projects(id, name, description, org_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllProjects :many
SELECT * FROM projects;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = ?;

-- name: GetProjectsByOrgID :many
SELECT * FROM projects
WHERE org_id = ?;