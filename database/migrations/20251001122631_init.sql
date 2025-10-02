-- +goose Up
CREATE TABLE auth (
    name VARCHAR(255),
    email VARCHAR(40) -- Comma removed from this line
);

-- +goose Down
DROP TABLE IF EXISTS auth;