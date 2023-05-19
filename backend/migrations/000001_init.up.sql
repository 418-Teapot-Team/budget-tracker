CREATE TABLE users
(
    id            bigint auto_increment primary key,
    full_name     varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null
);

create table accounts
(
    id               bigint auto_increment primary key,
    name             varchar(255)              NOT NULL,
    type             ENUM ('deposit','credit') NOT NULL,
    month_amount     tinyint                   NOT NULL,
    user_id          BIGINT                    NOT NULL,
    percent          FLOAT                     NOT NULL,
    date_opened      DATE                      NOT NULL,
    date_closed      DATE                      NOT NULL,
    monthly_payment  FLOAT                     not null,
    sum              FLOAT                     NOT NULL,
    created_at       DATETIME                  NOT NULL,
    already_received float                     not null,
    goal_sum         float                     not null,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE lists
(
    id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id     BIGINT                      NOT NULL,
    type        ENUM ('income', 'expenses') NOT NULL,
    category_id BIGINT                      NOT NULL,
    comment     VARCHAR(255)                NULL DEFAULT NULL,
    amount      FLOAT                       NOT NULL,
    created_at  DATETIME                    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

create table categories
(
    id         bigint auto_increment primary key,
    name       varchar(255) unique not null,
    type       varchar(255)        not null,
    color_hash varchar(255)        not null,
    image_link varchar(255)        not null
);
