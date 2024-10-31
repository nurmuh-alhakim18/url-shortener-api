-- name: CreateURL :one
INSERT INTO urls (custom_alias, original_url, expiration_date)
VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CheckCustomAlias :one
SELECT EXISTS (
  SELECT 1
  FROM urls
  WHERE custom_alias = $1
);

-- name: GetOriginalURL :one
SELECT original_url
FROM urls
WHERE custom_alias = $1 AND expiration_date > NOW();