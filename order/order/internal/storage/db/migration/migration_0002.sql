create table if not exists "order_items"
(
    id uuid
        constraint order_items_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    quantity int not null,
    price int not null,

    order_id uuid not null,

        constraint "fk_order" foreign key (order_id) references "order"(id) on delete cascade
);


