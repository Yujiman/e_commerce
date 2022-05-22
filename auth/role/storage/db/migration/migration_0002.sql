INSERT INTO role (id, created_at, updated_at, name, domain_id, scopes)
SELECT
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55',
    NOW(),
    NOW(),
    'user',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '{"role":"user"}'
WHERE NOT EXISTS (SELECT 1 FROM role WHERE name='user' and domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b');
