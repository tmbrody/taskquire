-- name: CreateTeam :execresult
INSERT INTO teams(id, name, description, owner_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllTeams :many
SELECT * FROM teams;

-- name: GetTeamByName :one
SELECT * FROM teams
WHERE name = ?;

-- name: GetTeamsByOwnerID :many
SELECT * FROM teams
WHERE owner_id = ?;

-- name: UpdateTeam :execresult
UPDATE teams
SET name = ?, description = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteTeam :execresult
DELETE FROM teams
WHERE id = ?;