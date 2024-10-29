-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id SERIAL PRIMARY KEY,
    short_url VARCHAR(10) NOT NULL UNIQUE ,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expiration_date TIMESTAMP,
    custom_alias VARCHAR(50) UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
