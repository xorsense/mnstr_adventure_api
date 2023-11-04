-- +goose Up
-- +goose StatementBegin
 CREATE TABLE IF NOT EXISTS adventurers (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL
 );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE adventurers;
-- +goose StatementEnd
