create table if not exists refresh_tokens
(
    id                uuid
        constraint refresh_tokens_pk
            primary key,
    access_token_id  uuid unique not null,
    expiry_date_time timestamp null
);
