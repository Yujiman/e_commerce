create table if not exists "category"
(
    id uuid
        constraint category_pk
            primary key,
    created_at          timestamp(6) NOT NULL,
    updated_at          timestamp(6) NOT NULL,
    group_id uuid not null ,
    "name" varchar(255) not null,

    constraint "fk_group" foreign key (group_id) references "group"(id) on delete cascade
    );


