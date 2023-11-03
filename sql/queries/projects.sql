-- name: CreateProject :execresult
INSERT INTO projects(id, name, description, org_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetProjectByName :one
SELECT * FROM projects
WHERE name = ?;

-- name: GetProjectByID :one
SELECT * FROM projects
WHERE id = ?;

-- name: GetProjectsByOrgID :many
SELECT * FROM projects
WHERE org_id = ?;

-- name: UpdateProject :execresult
UPDATE projects
SET name = ?, description = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteProject :execresult
DELETE FROM projects
WHERE id = ?;