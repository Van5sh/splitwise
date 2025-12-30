-- name: AddSettlement :one
INSERT INTO settlements (
    group_id,
    from_user_id,
    to_user_id,
    amount
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetSettlementById :one
SELECT * FROM settlements WHERE id = $1;

-- name: GetSettlementsByGroupId :many
SELECT * FROM settlements WHERE group_id = $1;

-- name: DeleteSettlementById :one
DELETE FROM settlements WHERE id=$1 RETURNING *;

-- name: ValidateUsersInSameGroup :one
SELECT 1
FROM user_groups
WHERE group_id = $3
    AND user_id IN ($1, $2)
GROUP BY group_id
HAVING COUNT(DISTINCT user_id) = 2;

-- name: GetSettlementsByUserId :many
SELECT * FROM settlements
WHERE from_user_id = $1 
    OR to_user_id = $1;