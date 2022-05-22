--  user16
INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '1d1501d1-0a2f-4661-b87a-2c7bf9816871',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user16',
    '$argon2id$v=19$m=65536,t=1,p=2$9TwhzOxg1dWPwt/+v07z5g$94eF12NvpsEbr1+JtwhPzforLMdICtpWdqx5CpTu53w',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user16');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '1d1501d1-0a2f-4661-b87a-2c7bf9816871',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='1d1501d1-0a2f-4661-b87a-2c7bf9816871'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);

-- user17
INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '3798d0c1-c8b7-4474-a918-8b484443145c',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user17',
    '$argon2id$v=19$m=65536,t=1,p=2$BpILE4EN8vd64yhVZbBR+Q$gr5pfd/dptYtYwLwIn6rEs1U9PVwQ7bfF0M6KRju8fU',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user17');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '3798d0c1-c8b7-4474-a918-8b484443145c',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='3798d0c1-c8b7-4474-a918-8b484443145c'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);

-- user18
INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'b72466b7-004c-4aab-ad25-47a88f2b8468',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user18',
    '$argon2id$v=19$m=65536,t=1,p=2$orXdhUU28I7/is/Jvo6XUA$tJuF8UFJPln02vJcPhgKqHlFYqSh4loCEuOvG85texM',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user18');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'b72466b7-004c-4aab-ad25-47a88f2b8468',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='b72466b7-004c-4aab-ad25-47a88f2b8468'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);

-- user19
INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'c633500a-a9b0-448a-8327-9e0d87cd8223',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user19',
    '$argon2id$v=19$m=65536,t=1,p=2$mQVvR3Vhz6TlExs//o+hTg$cdxYQLzPBwPFLadXI9IZ1TYCTFq+gFTZ/AMTqGIHZs4',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user19');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'c633500a-a9b0-448a-8327-9e0d87cd8223',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='c633500a-a9b0-448a-8327-9e0d87cd8223'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);

-- user20
INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '1e731fa4-7ff9-41f2-9762-7076bf9b50f0',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user20',
    '$argon2id$v=19$m=65536,t=1,p=2$zhgSKriMQ6nhiIlD3gHQew$VEp2Kh3u6Lrbu4vIBQoA2cABjDuCxF+mfJ/3ghPdfIQ',
    'active'
    WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user20');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '1e731fa4-7ff9-41f2-9762-7076bf9b50f0',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    WHERE NOT EXISTS (
    SELECT 1 FROM users_domains
    WHERE user_id='1e731fa4-7ff9-41f2-9762-7076bf9b50f0'
    AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
    AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);