-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    custom_alias VARCHAR(50) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expiration_date TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
