-- name: GetExpenses :many
SELECT * FROM expenses;

-- name: GetExpenseById :one
SELECT * FROM expenses WHERE id = $1;


-- name: CreateExpense :one
INSERT INTO expenses (
    group_id,
    paid_by,
    amount,
    description,
    created_at,
    paid
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: DeleteExpense :one
DELETE FROM expenses WHERE id = $1 
RETURNING *;

-- name: ValidateExpenseUserIsGroupMember :one
SELECT 1
FROM user_groups ug
WHERE ug.user_id = $1 AND ug.group_id = $2
LIMIT 1;

-- name: UpdateExpense :one
UPDATE expenses SET
    amount = COALESCE($2, amount),
    description = COALESCE($3, description),
    paid = $4,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: ValidateUserIsExpenseOwner :one
SELECT 1
FROM expenses
WHERE id = $1 AND paid_by = $2;

-- name: GetExpensesByGroupId :many
SELECT * FROM expenses WHERE group_id = $1;

-- name: GetExpensesByUserId :many
SELECT * FROM expenses WHERE paid_by = $1;

-- name: ValidateExpenseExists :one
SELECT 1 FROM expenses WHERE id = $1;
