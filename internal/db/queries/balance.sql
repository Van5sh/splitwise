-- name: GetUserBalanceInGroup :one
SELECT
  COALESCE(SUM(es.amount), 0)
  - COALESCE(SUM(CASE WHEN e.paid_by = $1 THEN e.amount END), 0) AS balance
FROM expenses e
LEFT JOIN expense_splits es
  ON e.id = es.expense_id AND es.user_id = $1
WHERE e.group_id = $2;

-- name: GetGroupBalance :many
SELECT
  u.id AS user_id,
  COALESCE(SUM(es.amount), 0)
  - COALESCE(SUM(CASE WHEN e.paid_by = u.id THEN e.amount END), 0) AS balance
FROM user_groups ug
JOIN users u ON u.id = ug.user_id
LEFT JOIN expenses e ON e.group_id = ug.group_id
LEFT JOIN expense_splits es ON e.id = es.expense_id AND es.user_id = u.id
WHERE ug.group_id = $1
GROUP BY u.id;


-- name: GetOverallUserBalance :one
SELECT
  COALESCE(SUM(es.amount), 0)
  - COALESCE(SUM(CASE WHEN e.paid_by = $1 THEN e.amount END), 0) AS balance
FROM expenses e
LEFT JOIN expense_splits es
  ON e.id = es.expense_id AND es.user_id = $1;
