// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: tasks.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createTask = `-- name: CreateTask :execresult
INSERT INTO tasks(id, name, description, project_id, team_id, owner_id, parent_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreateTaskParams struct {
	ID          string
	Name        string
	Description sql.NullString
	ProjectID   string
	TeamID      string
	OwnerID     string
	ParentID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.ProjectID,
		arg.TeamID,
		arg.OwnerID,
		arg.ParentID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteTask = `-- name: DeleteTask :execresult
DELETE FROM tasks
WHERE id = ?
`

func (q *Queries) DeleteTask(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTask, id)
}

const getSubtasksByParentID = `-- name: GetSubtasksByParentID :many
SELECT id, name, description, created_at, updated_at, project_id, team_id, owner_id, parent_id FROM tasks
WHERE parent_id = ?
`

func (q *Queries) GetSubtasksByParentID(ctx context.Context, parentID string) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getSubtasksByParentID, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
			&i.TeamID,
			&i.OwnerID,
			&i.ParentID,
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

const getTaskByID = `-- name: GetTaskByID :one
SELECT id, name, description, created_at, updated_at, project_id, team_id, owner_id, parent_id FROM tasks
WHERE id = ?
`

func (q *Queries) GetTaskByID(ctx context.Context, id string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByID, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
		&i.TeamID,
		&i.OwnerID,
		&i.ParentID,
	)
	return i, err
}

const getTaskByName = `-- name: GetTaskByName :one
SELECT id, name, description, created_at, updated_at, project_id, team_id, owner_id, parent_id FROM tasks
WHERE name = ?
`

func (q *Queries) GetTaskByName(ctx context.Context, name string) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTaskByName, name)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProjectID,
		&i.TeamID,
		&i.OwnerID,
		&i.ParentID,
	)
	return i, err
}

const getTasksByProjectID = `-- name: GetTasksByProjectID :many
SELECT id, name, description, created_at, updated_at, project_id, team_id, owner_id, parent_id FROM tasks
WHERE project_id = ?
`

func (q *Queries) GetTasksByProjectID(ctx context.Context, projectID string) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksByProjectID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
			&i.TeamID,
			&i.OwnerID,
			&i.ParentID,
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

const getTasksByTeamID = `-- name: GetTasksByTeamID :many
SELECT id, name, description, created_at, updated_at, project_id, team_id, owner_id, parent_id FROM tasks
WHERE team_id = ?
`

func (q *Queries) GetTasksByTeamID(ctx context.Context, teamID string) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getTasksByTeamID, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProjectID,
			&i.TeamID,
			&i.OwnerID,
			&i.ParentID,
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

const updateTask = `-- name: UpdateTask :execresult
UPDATE tasks
SET name = ?, description = ?, updated_at = ?
WHERE id = ?
`

type UpdateTaskParams struct {
	Name        string
	Description sql.NullString
	UpdatedAt   time.Time
	ID          string
}

func (q *Queries) UpdateTask(ctx context.Context, arg UpdateTaskParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateTask,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
}
