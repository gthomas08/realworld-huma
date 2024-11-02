-- +goose Up
-- +goose StatementBegin
ALTER TABLE tags
RENAME COLUMN tag TO name;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE tags
RENAME COLUMN name TO tag;

-- +goose StatementEnd