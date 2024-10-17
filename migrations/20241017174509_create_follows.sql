-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    follows (
        id UUID PRIMARY KEY,
        follower_id UUID REFERENCES users (id),
        followee_id UUID REFERENCES users (id)
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE follows;

-- +goose StatementEnd