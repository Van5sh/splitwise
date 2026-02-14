ALTER TABLE users
ADD COLUMN IF NOT EXISTS user_name TEXT;

UPDATE users u
SET user_name = ud.user_name
FROM user_details ud
WHERE ud.user_id = u.id
  AND u.user_name IS NULL;

ALTER TABLE users
ALTER COLUMN user_name SET NOT NULL;

ALTER TABLE users
DROP COLUMN IF EXISTS role;

ALTER TABLE user_details
DROP COLUMN IF EXISTS user_name;
