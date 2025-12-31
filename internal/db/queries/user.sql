-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

-- name: GetUserByFirebaseUid :one
SELECT * FROM users WHERE firebase_uid = $1;

-- name: GetUserByEmail :one
SELECT u.*, ud.user_name, ud.email
FROM users u
JOIN user_details ud ON u.id = ud.user_id
WHERE ud.email = $1;

-- name: CreateUser :one
INSERT INTO users (
    firebase_uid,
    role
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
    role = COALESCE($2, role),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: CheckUserExistsByEmail :one
SELECT 1 FROM user_details WHERE email = $1 LIMIT 1;

-- name: GetUserGroups :many
SELECT ug.*, g.group_name, g.description
FROM user_groups ug 
JOIN groups g ON ug.group_id = g.id
WHERE ug.user_id = $1;

-- name: GetUserSettlements :many
SELECT s.*
FROM settlements s
WHERE s.from_user_id = $1 OR s.to_user_id = $1;

-- name: GetUserExpensesByUserId :many
SELECT e.*
FROM expenses e
WHERE e.paid_by = $1;

-- name: ValidateUserExists :one
SELECT 1 FROM users WHERE id = $1;

-- name: GetUserWithDetails :one
SELECT u.*, ud.user_name, ud.email
FROM users u
JOIN user_details ud ON ud.user_id = u.id
WHERE u.id = $1;

