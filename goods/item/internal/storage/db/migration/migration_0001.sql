create table if not exists "item"
(
    id uuid
        constraint item_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
--  TODO Fill!
);
