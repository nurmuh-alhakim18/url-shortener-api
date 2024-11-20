-- name: CreateURL :one
INSERT INTO urls (custom_alias, original_url, expiration_date)
VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: CheckCustomAlias :one
SELECT EXISTS (
  SELECT 1
  FROM urls
  WHERE custom_alias = ?
);

-- name: GetOriginalURL :one
SELECT original_url
FROM urls
WHERE custom_alias = ? AND expiration_date > CURRENT_TIMESTAMP;