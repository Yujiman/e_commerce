create table if not exists "users"
(
    id uuid
    constraint users_user_pk
    primary key,
    created_at timestamp(6) not null,
    updated_at timestamp(6) not null,
    email varchar(255) unique,
    phone varchar(255) unique,
    login varchar(255) unique,
    password_hash varchar(255),
    status varchar(50) not null,
    role_id uuid not null
);
