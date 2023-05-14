CREATE TABLE users
(
    id            bigint auto_increment primary key,
    full_name     varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);