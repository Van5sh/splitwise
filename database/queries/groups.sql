-- name: GetGroups :many
SELECT *
FROM groups;


-- name: GetGroupById :one
SELECT *
FROM groups
WHERE id = $1;


-- name: GetGroupByName :one
SELECT *
FROM groups
WHERE group_name = $1;


-- name: GetGroupsByUserId :many
SELECT g.*
FROM groups g
JOIN user_groups ug ON g.id = ug.group_id
WHERE ug.user_id = $1;


-- name: CreateGroup :one
INSERT INTO groups (
    group_name,
    description
) VALUES (
    $1,
    $2
)
RETURNING *;


-- name: UpdateGroup :one
UPDATE groups SET
    group_name = COALESCE($2, group_name),
    description = COALESCE($3, description),
    updated_at = now()
WHERE id = $1
RETURNING *;


-- name: DeleteGroupById :one
DELETE FROM groups
WHERE id = $1
RETURNING *;


-- name: GetGroupMembers :many
SELECT u.*
FROM users u
JOIN user_groups ug ON u.id = ug.user_id
WHERE ug.group_id = $1;


-- name: GetGroupAdmins :many
SELECT u.*
FROM users u
JOIN user_groups ug ON u.id = ug.user_id
WHERE ug.group_id = $1
  AND ug.role = 'admin';


-- name: ValidateGroupExists :one
SELECT 1
FROM groups
WHERE id = $1;


-- name: CheckUserInGroup :one
SELECT 1
FROM user_groups
WHERE user_id = $1
  AND group_id = $2;


-- name: CheckUserIsGroupAdmin :one
SELECT 1
FROM user_groups
WHERE user_id = $1
  AND group_id = $2
  AND role = 'admin';

-- name: GetUserRoleInGroup :one
SELECT role
FROM user_groups
WHERE user_id = $1 AND group_id = $2;
