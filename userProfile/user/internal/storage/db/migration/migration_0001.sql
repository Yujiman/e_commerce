create table if not exists "user"
(
    id uuid
        constraint user_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    city_id uuid not null,
    phone varchar(11)not null ,
    firstname varchar(255) not null ,
    lastname varchar(255) not null ,
    patronymic varchar(255) not null
);
