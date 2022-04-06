create table if not exists "delivery_point"
(
    id uuid
        constraint delivery_point_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    city_id uuid not null,
    name varchar(255) not null ,
    address varchar(255) not null
);
