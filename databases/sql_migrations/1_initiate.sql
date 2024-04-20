-- +migrate Up
-- +migrate Down

CREATE TABLE person (
    id BIGINT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
);

-- +migrate StatementEnd