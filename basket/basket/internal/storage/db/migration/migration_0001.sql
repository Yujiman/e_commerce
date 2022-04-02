create table if not exists "basket"
(
    id uuid
        constraint basket_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    user_id uuid not null
);
