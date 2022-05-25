create table if not exists "category"
(
    id uuid
        constraint category_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    group_id uuid not null ,
    "name" varchar(255) not null,

    constraint "fk_group" foreign key (group_id) references "group"(id) on delete cascade
    );


INSERT INTO "category" (id, created_at, updated_at, group_id, name)
SELECT '47deec22-ac0d-430d-a9b8-3ac810db49f4', NOW(), NOW() ,'6afd70e0-2b6f-46c2-9292-287cbc25de5b', 'category_1'
WHERE NOT EXISTS (SELECT 1 FROM "category" WHERE id='47deec22-ac0d-430d-a9b8-3ac810db49f4');

INSERT INTO "category" (id, created_at, updated_at, group_id, name)
SELECT '16da1598-3181-4cbc-89ad-d564d4d368eb', NOW(), NOW() ,'6afd70e0-2b6f-46c2-9292-287cbc25de5b', 'category_2'
WHERE NOT EXISTS (SELECT 1 FROM "category" WHERE id='16da1598-3181-4cbc-89ad-d564d4d368eb');