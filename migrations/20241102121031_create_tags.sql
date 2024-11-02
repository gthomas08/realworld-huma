-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    tags (id UUID PRIMARY KEY, tag TEXT NOT NULL UNIQUE);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE tags;

-- +goose StatementEnd