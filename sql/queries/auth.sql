-- name: GetUserFromEmail :one
SELECT * FROM auth
WHERE email = $1 AND deleted = FALSE LIMIT 1;

-- name: GetUserFromId :one
SELECT * FROM auth
WHERE id = $1 AND deleted = FALSE LIMIT 1;

-- name: CreateUser :one
INSERT INTO auth (name, email,	password,	roles) 
VALUES ($1,	$2,	$3,	$4) RETURNING *;

-- name: UpdateUser :one
UPDATE auth SET
	name = $2,
	email = $3,
	password = $4,
	profile = $5,
	"emailVerified" = $6,
	roles = $7
WHERE id = $1 AND deleted = FALSE RETURNING *;

-- name: DeleteUserById :exec
UPDATE auth SET
	deleted = TRUE
WHERE id = $1 AND deleted = FALSE;

-- name: DeleteUserByEmail :exec
UPDATE auth SET
	deleted = TRUE
WHERE email = $1 AND deleted = FALSE;

-- name: UnDeleteUserById :exec
UPDATE auth SET
	deleted = FALSE
WHERE id = $1 AND deleted = TRUE;

-- name: UnDeleteUserByEmail :exec
UPDATE auth SET
	deleted = FALSE
WHERE email = $1 AND deleted = TRUE;

-- name: SearchUsersFromNameOrEmail :many
SELECT * FROM auth
WHERE name LIKE $1 OR email LIKE $1;

-- name: GetUserAllDetailsFromIdOrEmail :one
SELECT * FROM auth
INNER JOIN profile ON auth.profile = profile.id
INNER JOIN role ON auth.roles = role.id
INNER JOIN workspace ON role.workspace = workspace.id
WHERE auth.id = $1 OR auth.email = $1
LIMIT 1;

-- name: GetProfileFromId :one
SELECT * FROM profile
WHERE id = $1 LIMIT 1;

-- name: GetProfileFromUserId :one
SELECT * FROM profile
WHERE "authId" = $1 LIMIT 1;

-- name: GetWorkspaceFromName :one
SELECT * FROM workspace
WHERE name = $1 LIMIT 1;

-- name: GetRoleFromName :one
SELECT * FROM role
WHERE name = $1 LIMIT 1;

-- name: GetPaginatedWorkspaces :many
SELECT * FROM workspace
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetPaginatedUsers :many
SELECT * FROM auth
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetPaginatedRoles :many
SELECT * FROM role
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: GetPaginatedPopulatedUsers :many
SELECT auth.*, role.name FROM auth
INNER JOIN role ON auth.roles = role.id
ORDER BY auth.id
LIMIT $1 OFFSET $2;

-- name: GetPaginatedPopulatedRoles :many
SELECT role.*, workspace.name FROM role
INNER JOIN workspace ON role.workspace = workspace.id
ORDER BY role.id
LIMIT $1 OFFSET $2;
