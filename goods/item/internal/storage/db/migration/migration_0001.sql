create table if not exists "group"
(
    id uuid
        constraint group_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    "name" varchar(255) not null
);