-- name: AddExpenseSplits :one
INSERT INTO expense_splits (
    expense_id,
    user_id,
    amount
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetSplitsByExpenseId :many
SELECT * FROM expense_splits 
WHERE expense_id = $1;

-- name: GetSplitsByUserId :many
SELECT * FROM expense_splits 
WHERE user_id = $1;

-- name: ValidateSplitTotalEqualsExpense :one
SELECT 1
FROM expenses e
WHERE e.id = $1 AND e.amount = (
    SELECT COALESCE(SUM(es.amount), 0)
    FROM expense_splits es
    WHERE es.expense_id = e.id
);
