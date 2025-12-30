CREATE TABLE user_groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id UUID NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
    role TEXT NOT NULL DEFAULT 'member'
        CHECK (role IN ('member', 'admin')),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (user_id, group_id)
);

CREATE UNIQUE INDEX one_admin_per_group
ON user_groups (group_id)
WHERE role = 'admin';

CREATE INDEX idx_user_groups_user
ON user_groups (user_id);

CREATE INDEX idx_user_groups_group
ON user_groups (group_id);