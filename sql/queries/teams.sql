-- name: CreateTeam :execresult
INSERT INTO teams(id, name, created_at, updated_at, org_id)
VALUES (?, ?, ?, ?, ?);

-- name: GetAllTeams :many
SELECT * FROM teams;