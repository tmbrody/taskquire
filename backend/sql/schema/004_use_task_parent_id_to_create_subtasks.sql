-- +goose Up
ALTER TABLE tasks ADD COLUMN parent_id TEXT NOT NULL REFERENCES tasks(id);

-- +goose Down
ALTER TABLE tasks DROP COLUMN parent_id;