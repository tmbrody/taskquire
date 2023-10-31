-- name: CreateOrg :execresult
INSERT INTO orgs(id, name, description, owner_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAllOrgs :many
SELECT * FROM orgs;

-- name: GetOrgByName :one
SELECT * FROM orgs
WHERE name = ?;

-- name: UpdateOrg :execresult
UPDATE orgs
SET name = ?, description = ?, updated_at = ?
WHERE id = ?;

-- name: DeleteOrg :execresult
DELETE FROM orgs
WHERE id = ?;