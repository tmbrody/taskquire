// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: projects.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createProject = `-- name: CreateProject :execresult
INSERT INTO projects(id, name, description, org_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateProjectParams struct {
	ID          string
	Name        string
	Description string
	OrgID       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateProject(ctx context.Context, arg CreateProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createProject,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.OrgID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteProject = `-- name: DeleteProject :execresult
DELETE FROM projects
WHERE id = ?
`

func (q *Queries) DeleteProject(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteProject, id)
}

const getProjectByID = `-- name: GetProjectByID :one
SELECT id, name, description, created_at, updated_at, org_id FROM projects
WHERE id = ?
`

func (q *Queries) GetProjectByID(ctx context.Context, id string) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByID, id)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrgID,
	)
	return i, err
}

const getProjectByName = `-- name: GetProjectByName :one
SELECT id, name, description, created_at, updated_at, org_id FROM projects
WHERE name = ?
`

func (q *Queries) GetProjectByName(ctx context.Context, name string) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByName, name)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrgID,
	)
	return i, err
}

const getProjectsByOrgID = `-- name: GetProjectsByOrgID :many
SELECT id, name, description, created_at, updated_at, org_id FROM projects
WHERE org_id = ?
`

func (q *Queries) GetProjectsByOrgID(ctx context.Context, orgID string) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjectsByOrgID, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
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

const updateProject = `-- name: UpdateProject :execresult
UPDATE projects
SET name = ?, description = ?, updated_at = ?
WHERE id = ?
`

type UpdateProjectParams struct {
	Name        string
	Description string
	UpdatedAt   time.Time
	ID          string
}

func (q *Queries) UpdateProject(ctx context.Context, arg UpdateProjectParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateProject,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
}
