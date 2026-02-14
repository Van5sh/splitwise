-- name: GetUserDetailsByUserId :one
SELECT user_details.id,
       user_details.user_id,
       user_details.email,
       user_details.created_at,
       user_details.updated_at,
       users.id,
       users.firebase_uid,
       users.user_name,
       users.created_at,
       users.updated_at
FROM user_details
JOIN users ON user_details.user_id = users.id
WHERE user_details.user_id = $1;

-- name: GetUserDetailsByEmail :one
SELECT *
FROM user_details
WHERE email = $1;

-- name: CreateUserDetails :one
INSERT INTO user_details (user_id, email)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUserDetails :one
UPDATE user_details
SET email = COALESCE($2, email),
    updated_at = now()
WHERE user_id = $1
RETURNING *;

-- name: UpdateUserName :one
UPDATE users
SET user_name = COALESCE($2, user_name),
    updated_at = now()
WHERE id = $1
RETURNING *;

