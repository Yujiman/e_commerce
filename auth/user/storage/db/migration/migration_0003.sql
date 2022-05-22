create table if not exists users_domains
(
    user_id   uuid not null,
    domain_id uuid not null,
    role_id uuid not null,

    PRIMARY KEY (user_id, domain_id),
    constraint fk_user foreign key (user_id) references users_user(id) on delete cascade,
    constraint fk_domain foreign key (domain_id) references domain(id) on delete cascade
);