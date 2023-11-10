-- name: AddTeamToProject :execresult
INSERT INTO projects_teams(project_id, team_id)
VALUES (?, ?);

-- name: GetOneProjectFromTeam :one
SELECT * FROM projects_teams
WHERE project_id = ?
AND team_id = ?;

-- name: GetAllTeamsFromProject :many
SELECT * FROM projects_teams
WHERE project_id = ?;

-- name: GetAllProjectsByTeam :many
SELECT * FROM projects_teams
WHERE team_id = ?;

-- name: RemoveProjectFromTeam :execresult
DELETE FROM projects_teams
WHERE project_id = ?
AND team_id = ?;