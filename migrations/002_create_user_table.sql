-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    participant_id INTEGER PRIMARY KEY REFERENCES participant(id) ON DELETE CASCADE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    middle_name VARCHAR(100),
    city VARCHAR(100),
    age INTEGER,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users"
-- +goose StatementEnd
