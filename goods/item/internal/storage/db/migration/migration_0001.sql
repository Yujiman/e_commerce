create table if not exists "group"
(
    id uuid
        constraint group_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    "name" varchar(255) not null
);

INSERT INTO "group" (id, created_at, updated_at, name)
SELECT '6afd70e0-2b6f-46c2-9292-287cbc25de5b', NOW(), NOW() ,'group1'
WHERE NOT EXISTS (SELECT 1 FROM "group" WHERE id='6afd70e0-2b6f-46c2-9292-287cbc25de5b');
