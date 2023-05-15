CREATE TABLE users
(
    id            bigint auto_increment primary key,
    full_name     varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table lists
(
    id         bigint auto_increment primary key,
    user_id    bigint                    not null,
    type       ENUM ('income','outcome') not null,
    category   varchar(255)              not null,
    comment    varchar(255)              null default null,
    amount     float                     not null,
    created_at datetime                  not null,
    foreign key (user_id) references users (id)
);