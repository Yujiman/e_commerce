create table if not exists "item"
(
    id uuid
    constraint item_pk
    primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    category_id uuid not null ,
    brand varchar(255) not null,
    "name" varchar(255) not null,
    description varchar(255) not null ,
    image_link varchar (255) not null ,
    price decimal not null,

    constraint fk_category foreign key (category_id) references category(id) on delete cascade
    );