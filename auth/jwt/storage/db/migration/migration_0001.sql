create table if not exists access_tokens
(
    id                uuid
        constraint access_tokens_pk
            primary key,
    expiry_date_time  timestamp    not null,
    user_identifier   char(36)     not null,
    domain_identifier char(36)     not null,
    client            varchar(255) not null,
    scopes            text         not null
);

create index if not exists access_tokens_user_domain_index
    on "access_tokens" (user_identifier, domain_identifier);
