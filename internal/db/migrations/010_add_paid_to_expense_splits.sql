ALTER TABLE expense_splits
ADD COLUMN paid_to UUID;

UPDATE expense_splits es
SET paid_to = e.paid_by
FROM expenses e
WHERE e.id = es.expense_id
  AND es.paid_to IS NULL;

ALTER TABLE expense_splits
ALTER COLUMN paid_to SET NOT NULL;

ALTER TABLE expense_splits
ADD CONSTRAINT expense_splits_paid_to_fkey
FOREIGN KEY (paid_to) REFERENCES users(id);

CREATE INDEX idx_expense_splits_paid_to
ON expense_splits (paid_to);
