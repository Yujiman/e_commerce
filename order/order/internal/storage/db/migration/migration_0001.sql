create table if not exists "order"
(
    id uuid
        constraint order_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    client_id uuid not null,
    status varchar(255) not null ,
    order_number serial,
    is_payed bool
);
