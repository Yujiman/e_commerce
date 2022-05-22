INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '1d1f550f-2a0b-48a4-92b1-f4f1461bba44',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user11',
    '$argon2id$v=19$m=65536,t=1,p=2$rhsBSlD+Ro1ptRpTfny9EQ$B7qfAQ/bHJGIW9lI7tXTwXM69giU6iXPRSMq9QurwEY',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user11');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '1d1f550f-2a0b-48a4-92b1-f4f1461bba44',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='1d1f550f-2a0b-48a4-92b1-f4f1461bba44'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '3c1b2a39-7e1a-4664-9b26-696734c51b64',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user12',
    '$argon2id$v=19$m=65536,t=1,p=2$JEwG8QvT2YUPuB0gikNd4g$mquIpeyeAA4MpPo64qYFJP+k2P9KgjFgxbGEEYg70dc',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user12');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '3c1b2a39-7e1a-4664-9b26-696734c51b64',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='3c1b2a39-7e1a-4664-9b26-696734c51b64'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'fbbe7966-2a3e-4108-a017-6e450c5e1391',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user13',
    '$argon2id$v=19$m=65536,t=1,p=2$uGLk2LYl9dSxKy0R7W30Xw$yBHd8kLzRAyQ0U6oKqk6U7PHJNETe1SzG/EiktXwU3c',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user13');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'fbbe7966-2a3e-4108-a017-6e450c5e1391',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='fbbe7966-2a3e-4108-a017-6e450c5e1391'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );



INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    'a6add85f-5d21-4dd6-880d-d12ae3e51b03',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user14',
    '$argon2id$v=19$m=65536,t=1,p=2$zDbrR2sf1ZVBP2RlIGCotw$uhBoZYr2v9LjbuQOlQ87sp0FAIBDT37ZuBjHixFTq/g',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user14');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    'a6add85f-5d21-4dd6-880d-d12ae3e51b03',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
        WHERE user_id='a6add85f-5d21-4dd6-880d-d12ae3e51b03'
          AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
          AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
    );

INSERT INTO users_user (id, created_at, updated_at, email, phone, login, password_hash, status)
SELECT
    '2aaf7b5d-092b-41a6-bd36-ec2fe5928efe',
    NOW(),
    NOW(),
    NULL,
    NULL,
    'user15',
    '$argon2id$v=19$m=65536,t=1,p=2$/EoHvJZdvLPKE4Jmipkgxw$kktHo0qDFVRPgUctEF9a50rCJqCiIXybbXw2TNGdaBk',
    'active'
WHERE NOT EXISTS (SELECT 1 FROM users_user WHERE login='user15');

INSERT INTO users_domains (user_id, domain_id, role_id)
SELECT
    '2aaf7b5d-092b-41a6-bd36-ec2fe5928efe',
    '6afd70e0-2b6f-46c2-9292-287cbc25de5b',
    '86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
WHERE NOT EXISTS (
        SELECT 1 FROM users_domains
    WHERE user_id='2aaf7b5d-092b-41a6-bd36-ec2fe5928efe'
      AND domain_id='6afd70e0-2b6f-46c2-9292-287cbc25de5b'
      AND role_id='86bdd3e1-5f9c-4898-bf57-4c74d8208b55'
);
