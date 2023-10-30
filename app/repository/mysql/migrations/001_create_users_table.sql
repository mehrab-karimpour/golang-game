-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id           int(11)      NOT NULL AUTO_INCREMENT,
    phone_number varchar(15)  NOT NULL UNIQUE,
    first_name   varchar(100)          DEFAULT NULL,
    last_name    varchar(100)          DEFAULT NULL,
    password     varchar(255) NOT NULL,
    created_at   timestamp    NOT NULL DEFAULT current_timestamp(),
    updated_at   timestamp    NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

-- +migrate Down

DROP TABLE users;