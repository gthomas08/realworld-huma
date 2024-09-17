-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id UUID PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        bio TEXT,
        image TEXT
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

-- +goose StatementEnd