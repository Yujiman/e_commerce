create table if not exists "user"
(
    id uuid
        constraint user_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    city_id uuid not null,
    phone varchar(11)not null ,
    first_name varchar(255) not null ,
    last_name varchar(255) not null ,
    middle_name varchar(255) not null
);
