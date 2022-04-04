create table if not exists "delivery_point_user"
(
    user_id uuid
        constraint delivery_point_user_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    delivery_point_id uuid not null
);
