INSERT INTO domain (id, name, url, created_at, updated_at)
SELECT '6afd70e0-2b6f-46c2-9292-287cbc25de5b', 'gcc-web', 'gcc.com', NOW(), NOW()
    WHERE NOT EXISTS (SELECT 1 FROM domain WHERE url='gcc.com');

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'af96043d-5703-45fb-9832-c7da4cd2c753',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user1',
    '$argon2id$v=19$m=65536,t=1,p=2$imykSTJUG5fI/6fP+oTxmw$5XkymQse0PnXl6neYtAKPvjliKn3FCuRSStnJsCitBE',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user1');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'af96043d-5703-45fb-9832-c7da4cd2c753',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='af96043d-5703-45fb-9832-c7da4cd2c753'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '0f2a0da9-2c3e-4711-8e7c-d6c7cabe8547',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user2',
    '$argon2id$v=19$m=65536,t=1,p=2$imykSTJUG5fI/6fP+oTxmw$5XkymQse0PnXl6neYtAKPvjliKn3FCuRSStnJsCitBE',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user2');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '0f2a0da9-2c3e-4711-8e7c-d6c7cabe8547',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='0f2a0da9-2c3e-4711-8e7c-d6c7cabe8547'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);