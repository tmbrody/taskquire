// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: teams.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createTeam = `-- name: CreateTeam :execresult
INSERT INTO teams(id, name, created_at, updated_at, org_id)
VALUES (?, ?, ?, ?, ?)
`

type CreateTeamParams struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	OrgID     string
}

func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTeam,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.OrgID,
	)
}

const getAllTeams = `-- name: GetAllTeams :many
SELECT id, name, created_at, updated_at, org_id FROM teams
`

func (q *Queries) GetAllTeams(ctx context.Context) ([]Team, error) {
	rows, err := q.db.QueryContext(ctx, getAllTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Team
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OrgID,
		); err != nil {
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

const getTeamById = `-- name: GetTeamById :one
SELECT id, name, created_at, updated_at, org_id FROM teams
WHERE id = ?
`

func (q *Queries) GetTeamById(ctx context.Context, id string) (Team, error) {
	row := q.db.QueryRowContext(ctx, getTeamById, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrgID,
	)
	return i, err
}
