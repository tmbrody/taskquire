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
INSERT INTO tasks(id, name, description, created_at, updated_at, project_id)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateTaskParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProjectID   string
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTask,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.ProjectID,
	)
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, name, description, created_at, updated_at, project_id FROM tasks
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAllTasks)
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