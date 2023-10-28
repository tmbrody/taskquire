-- +goose Up
ALTER TABLE orgs
    ADD COLUMN description TEXT NOT NULL,
    ADD COLUMN creator_id TEXT NOT NULL;

ALTER TABLE projects
    ADD COLUMN org_id TEXT NOT NULL REFERENCES orgs(id);

CREATE TABLE projects_teams (
    project_id CHAR(36) NOT NULL REFERENCES projects(id),
    team_id CHAR(36) NOT NULL REFERENCES teams(id),
    UNIQUE(project_id, team_id)
);

CREATE TABLE teams_users (
    team_id CHAR(36) NOT NULL REFERENCES teams(id),
    user_id CHAR(36) NOT NULL REFERENCES users(id),
    UNIQUE(team_id, user_id)
);

-- +goose Down
ALTER TABLE orgs
    DROP COLUMN description,
    DROP COLUMN creator_id;

ALTER TABLE projects
    DROP COLUMN org_id;

DROP TABLE projects_teams;

DROP TABLE teams_users;