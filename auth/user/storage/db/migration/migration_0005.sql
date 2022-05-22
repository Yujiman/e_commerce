INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'dd7d2914-6e8f-4fbe-9872-82d447f7091b',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user3',
    '$argon2id$v=19$m=65536,t=1,p=2$R/knJizMdtlBsPc2N06XmQ$BzKSff2AeAgxfe2+2edB0lylFINrPIq8cb+Bh8PC6HM',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user3');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'dd7d2914-6e8f-4fbe-9872-82d447f7091b',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='dd7d2914-6e8f-4fbe-9872-82d447f7091b'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '1ddcd76f-3eb3-49f5-ab63-80a62d6665eb',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user4',
    '$argon2id$v=19$m=65536,t=1,p=2$pSbMAdMeRxBSXGSnUb3ZRw$OjRWiYv3H+DKKW1E8pW4xCC40XTBkT6viI0loxz1/mk',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user4');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '1ddcd76f-3eb3-49f5-ab63-80a62d6665eb',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='1ddcd76f-3eb3-49f5-ab63-80a62d6665eb'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '97af6117-b152-451d-86e7-d9bd3047ea26',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user5',
    '$argon2id$v=19$m=65536,t=1,p=2$Q7N+zkJw0NSVjWqfL9M8Vw$Tjrwog8rzbggpvEHluCIO03CxmfNqo/JiWxS4cafLpY',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user5');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '97af6117-b152-451d-86e7-d9bd3047ea26',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='97af6117-b152-451d-86e7-d9bd3047ea26'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '76866b9c-5e49-43de-b325-8ecf81e7e1c3',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user6',
    '$argon2id$v=19$m=65536,t=1,p=2$11Xv9oYCExTtpZXTjWvDUg$5P/NS1nSTifS2jEmTQxLgq7fXIQN7VhKuXGihYE+9e0',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user6');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '76866b9c-5e49-43de-b325-8ecf81e7e1c3',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='76866b9c-5e49-43de-b325-8ecf81e7e1c3'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '0f64075e-0a4b-4493-9f83-e158ca0be1fb',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user7',
    '$argon2id$v=19$m=65536,t=1,p=2$18mQFv5rRtOC/PC4sTs0yg$xvIHU+UH/sK/T9ZUYpk29scb772tiYYgnuLmKuq0G7A',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user7');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '0f64075e-0a4b-4493-9f83-e158ca0be1fb',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='0f64075e-0a4b-4493-9f83-e158ca0be1fb'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '610c18dd-91d2-4d51-968f-7d81a18849e7',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user8',
    '$argon2id$v=19$m=65536,t=1,p=2$TQcw3D+dFmGOT6q+GW0G8w$Xv0PnGris7ZurwwsE1QElVSeUwHYKPNJMCzVaZeYqPQ',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user8');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '610c18dd-91d2-4d51-968f-7d81a18849e7',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='610c18dd-91d2-4d51-968f-7d81a18849e7'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'dab249b6-5561-4951-a031-b1f13feefca4',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user9',
    '$argon2id$v=19$m=65536,t=1,p=2$2BlqJ1PQg4KtnUjHJwq3Bg$ed6++lOadwgIYaaOtI68BmtZXMgMi+EDB/HGCMS1QY8',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user9');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'dab249b6-5561-4951-a031-b1f13feefca4',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='dab249b6-5561-4951-a031-b1f13feefca4'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '0f3b4e49-ee24-4788-a6bf-d0e35e224490',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user10',
    '$argon2id$v=19$m=65536,t=1,p=2$XvERPsROCm4vmrudiAY3RQ$73+ziexu8cJhJE6jMeryrIqRB6nT9KrNRdLeSv1Pt1E',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user10');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '0f3b4e49-ee24-4788-a6bf-d0e35e224490',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='0f3b4e49-ee24-4788-a6bf-d0e35e224490'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );
