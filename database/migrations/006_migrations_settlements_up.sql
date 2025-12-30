CREATE TABLE settlements (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    from_user_id UUID NOT NULL REFERENCES users(id),
    to_user_id   UUID NOT NULL REFERENCES users(id),
    amount NUMERIC(12,2) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    CHECK (from_user_id <> to_user_id)
);

CREATE INDEX idx_settlements_group
ON settlements (group_id);

CREATE INDEX idx_settlements_from_user
ON settlements (from_user_id);

CREATE INDEX idx_settlements_to_user
ON settlements (to_user_id);