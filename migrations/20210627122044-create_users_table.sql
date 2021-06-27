
-- +migrate Up
CREATE TABLE users
(
    id         SERIAL,
    uid        VARCHAR(128)                          NOT NULL,
    username   VARCHAR(255)                          NOT NULL,
    email      VARCHAR(255)                          NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX users_uid_uindex ON users (uid);

-- +migrate Down
DROP TABLE users;