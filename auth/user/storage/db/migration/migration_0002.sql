create table if not exists domain
(
    id         uuid         not null
        constraint domain_pk
            primary key,
    name       varchar(255) unique not null,
    url        varchar(255) unique not null,
    created_at timestamp(6) not null,
    updated_at timestamp(6) not null
);
