CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    group_name TEXT NOT NULL,
    description TEXT,
    total_amount NUMERIC(12,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);
