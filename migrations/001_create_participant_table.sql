-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "participant";
-- +goose StatementEnd
