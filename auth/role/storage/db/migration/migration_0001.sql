create table if not exists "role"
(
    id uuid
    constraint role_pk
    primary key,
    created_at timestamp(6) not null,
    updated_at timestamp(6) not null,
    name varchar(100) not null,
    domain_id uuid not null,
    scopes text not null
);

create unique index if not exists role_name_domain_idx ON role(name, domain_id);
