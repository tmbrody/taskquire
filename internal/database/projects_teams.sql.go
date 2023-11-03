// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: projects_teams.sql

package database

import (
	"context"
	"database/sql"
)

const addTeamToProject = `-- name: AddTeamToProject :execresult
INSERT INTO projects_teams(project_id, team_id)
VALUES (?, ?)
`

type AddTeamToProjectParams struct {
	ProjectID string
	TeamID    string
}

func (q *Queries) AddTeamToProject(ctx context.Context, arg AddTeamToProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTeamToProject, arg.ProjectID, arg.TeamID)
}

const getAllProjectsByTeam = `-- name: GetAllProjectsByTeam :many
SELECT project_id, team_id FROM projects_teams
WHERE team_id = ?
`

func (q *Queries) GetAllProjectsByTeam(ctx context.Context, teamID string) ([]ProjectsTeam, error) {
	rows, err := q.db.QueryContext(ctx, getAllProjectsByTeam, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectsTeam
	for rows.Next() {
		var i ProjectsTeam
		if err := rows.Scan(&i.ProjectID, &i.TeamID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTeamsFromProject = `-- name: GetAllTeamsFromProject :many
SELECT project_id, team_id FROM projects_teams
WHERE project_id = ?
`

func (q *Queries) GetAllTeamsFromProject(ctx context.Context, projectID string) ([]ProjectsTeam, error) {
	rows, err := q.db.QueryContext(ctx, getAllTeamsFromProject, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectsTeam
	for rows.Next() {
		var i ProjectsTeam
		if err := rows.Scan(&i.ProjectID, &i.TeamID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOneProjectFromTeam = `-- name: GetOneProjectFromTeam :one
SELECT project_id, team_id FROM projects_teams
WHERE project_id = ?
AND team_id = ?
`

type GetOneProjectFromTeamParams struct {
	ProjectID string
	TeamID    string
}

func (q *Queries) GetOneProjectFromTeam(ctx context.Context, arg GetOneProjectFromTeamParams) (ProjectsTeam, error) {
	row := q.db.QueryRowContext(ctx, getOneProjectFromTeam, arg.ProjectID, arg.TeamID)
	var i ProjectsTeam
	err := row.Scan(&i.ProjectID, &i.TeamID)
	return i, err
}

const removeProjectFromTeam = `-- name: RemoveProjectFromTeam :execresult
DELETE FROM projects_teams
WHERE project_id = ?
AND team_id = ?
`

type RemoveProjectFromTeamParams struct {
	ProjectID string
	TeamID    string
}

func (q *Queries) RemoveProjectFromTeam(ctx context.Context, arg RemoveProjectFromTeamParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, removeProjectFromTeam, arg.ProjectID, arg.TeamID)
}
