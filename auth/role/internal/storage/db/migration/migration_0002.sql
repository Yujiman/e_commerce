INSERT INTO "role" (id, created_at, updated_at, name, scopes)
SELECT
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55',
    NOW(),
    NOW(),
    'user',
    '{"role":"user"}'
WHERE NOT EXISTS (SELECT 1 FROM "role" WHERE name='user' );
