-- name: CreateURL :one
INSERT INTO urls (short_url, original_url, expiration_date, custom_alias)
VALUES (
  $1, $2, $3, $4
)
RETURNING *;