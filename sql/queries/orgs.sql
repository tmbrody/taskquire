-- name: CreateOrg :execresult
INSERT INTO orgs(id, name, created_at, updated_at)
VALUES (?, ?, ?, ?);

-- name: GetAllOrgs :many
SELECT * FROM orgs;