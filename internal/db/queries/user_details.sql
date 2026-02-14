-- name: GetUserDetailsByUserId :one
SELECT *
FROM user_details
JOIN users ON user_details.user_id = users.id
WHERE user_details.user_id = $1;

-- name: GetUserDetailsByEmail :one
SELECT *
FROM user_details
WHERE email = $1;

-- name: CreateUserDetails :one
INSERT INTO user_details (user_id, user_name, email)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUserDetails :one
UPDATE user_details
SET user_name = COALESCE($2, user_name),
    email = COALESCE($3, email),
    updated_at = now()
WHERE user_id = $1
RETURNING *;


