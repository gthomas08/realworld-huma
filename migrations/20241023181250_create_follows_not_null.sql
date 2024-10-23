-- +goose Up
-- +goose StatementBegin
ALTER TABLE follows
ALTER COLUMN follower_id
SET
    NOT NULL,
ALTER COLUMN followee_id
SET
    NOT NULL;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE follows
ALTER COLUMN follower_id
DROP NOT NULL,
ALTER COLUMN followee_id
DROP NOT NULL;

-- +goose StatementEnd