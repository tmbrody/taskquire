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

-- name: GetAllProjectsFromTeam :many
SELECT * FROM projects_teams
WHERE team_id = ?;

-- name: RemoveTeamFromProject :execresult
DELETE FROM projects_teams
WHERE team_id = ?
AND project_id = ?;