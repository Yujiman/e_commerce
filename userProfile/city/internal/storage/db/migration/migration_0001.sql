create table if not exists "city"
(
    id uuid
        constraint city_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    name_ru varchar(255) not null,
    name_en varchar(255) not null
);
