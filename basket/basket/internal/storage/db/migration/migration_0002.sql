create table if not exists "basket_item"
(
    id uuid
        constraint basket_item_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    basket_id uuid not null ,
    price int8 not null,
    good_id uuid not null,
    quantity int8,

    constraint fk_basket foreign key (basket_id) references "basket"(id) on delete cascade

);
