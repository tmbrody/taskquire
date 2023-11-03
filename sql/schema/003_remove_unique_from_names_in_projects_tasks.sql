-- +goose Up
ALTER TABLE projects
    DROP INDEX name;

ALTER TABLE tasks
    ADD COLUMN team_id TEXT NOT NULL REFERENCES teams(id),
    ADD COLUMN owner_id TEXT NOT NULL REFERENCES users(id),
    DROP INDEX name;

-- +goose Down
ALTER TABLE projects
    ADD UNIQUE(name(255));

ALTER TABLE tasks
    DROP COLUMN team_id,
    DROP COLUMN owner_id,
    ADD UNIQUE(name(255));
