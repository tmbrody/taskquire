-- name: AddUsertoTeam :execresult
INSERT INTO teams_users(user_id, team_id)
VALUES (?, ?);

-- name: GetOneUserFromTeam :one
SELECT * FROM teams_users
WHERE user_id = ?
AND team_id = ?;

-- name: GetAllUsersFromTeam :many
SELECT * FROM teams_users
WHERE team_id = ?;

-- name: GetAllTeamsByUser :many
SELECT * FROM teams_users
WHERE user_id = ?;

-- name: RemoveUserFromTeam :execresult
DELETE FROM teams_users
WHERE user_id = ?
AND team_id = ?;