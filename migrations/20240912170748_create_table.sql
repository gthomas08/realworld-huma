-- +goose Up
-- +goose StatementBegin
CREATE TABLE pings (
    id INTEGER PRIMARY KEY,
    message TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pings;
-- +goose StatementEnd
