-- name: GetUserBalanceInGroup :one
SELECT
  COALESCE(SUM(es.amount), 0)
  - COALESCE(SUM(e.amount), 0) AS balance
FROM expenses e
LEFT JOIN expense_splits es
  ON e.id = es.expense_id AND es.user_id = $1
WHERE e.group_id = $2;

-- name: GetGroupBalance :many
SELECT
  es.user_id,
  SUM(es.amount) - SUM(e.amount) AS balance
FROM expenses e
JOIN expense_splits es ON e.id = es.expense_id
WHERE e.group_id = $1
GROUP BY es.user_id;


-- name: GetOverallUserBalance :one
SELECT
  COALESCE(SUM(es.amount), 0)
  - COALESCE(SUM(e.amount), 0) AS balance
FROM expenses e
LEFT JOIN expense_splits es
  ON e.id = es.expense_id AND es.user_id = $1;
