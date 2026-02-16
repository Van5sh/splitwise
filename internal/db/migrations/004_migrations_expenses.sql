CREATE TABLE expenses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    paid_by UUID NOT NULL REFERENCES users(id),
    description TEXT,
    amount NUMERIC(12,2) NOT NULL,
    paid BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_expenses_group
ON expenses (group_id);

CREATE INDEX idx_expenses_paid_by
ON expenses (paid_by);

CREATE INDEX idx_expenses_paid
ON expenses (paid);
