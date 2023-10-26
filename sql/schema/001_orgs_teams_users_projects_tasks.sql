-- +goose Up
CREATE TABLE users (
    id          CHAR(36)    PRIMARY KEY,
    name        TEXT        NOT NULL,
    email       TEXT        NOT NULL,
    password    TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    UNIQUE(email(255))
);

CREATE TABLE orgs (
    id          CHAR(36)    PRIMARY KEY,
    name        TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    UNIQUE(name(255))
);

CREATE TABLE teams(
    id          CHAR(36)    PRIMARY KEY,
    name        TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    org_id      CHAR(36)    NOT NULL    REFERENCES orgs(id),
    UNIQUE(name(255))
);

CREATE TABLE projects (
    id          CHAR(36)    PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    UNIQUE(name(255))
);

CREATE TABLE tasks (
    id          CHAR(36)    PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT,
    created_at  TIMESTAMP   NOT NULL,
    updated_at  TIMESTAMP   NOT NULL,
    project_id  CHAR(36)    NOT NULL    REFERENCES projects(id),
    UNIQUE(name(255))
);

-- +goose Down
DROP TABLE tasks;
DROP TABLE projects;
DROP TABLE users;
DROP TABLE teams;
DROP TABLE orgs;