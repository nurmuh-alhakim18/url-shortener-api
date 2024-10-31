// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: urls.sql

package repositories

import (
	"context"
	"database/sql"
)

const checkCustomAlias = `-- name: CheckCustomAlias :one
SELECT EXISTS (
  SELECT 1
  FROM urls
  WHERE custom_alias = $1
)
`

func (q *Queries) CheckCustomAlias(ctx context.Context, customAlias string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkCustomAlias, customAlias)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createURL = `-- name: CreateURL :one
INSERT INTO urls (custom_alias, original_url, expiration_date)
VALUES (
  $1, $2, $3
)
RETURNING id, custom_alias, original_url, created_at, expiration_date
`

type CreateURLParams struct {
	CustomAlias    string
	OriginalUrl    string
	ExpirationDate sql.NullTime
}

func (q *Queries) CreateURL(ctx context.Context, arg CreateURLParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, createURL, arg.CustomAlias, arg.OriginalUrl, arg.ExpirationDate)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.CustomAlias,
		&i.OriginalUrl,
		&i.CreatedAt,
		&i.ExpirationDate,
	)
	return i, err
}

const getOriginalURL = `-- name: GetOriginalURL :one
SELECT original_url
FROM urls
WHERE custom_alias = $1 AND expiration_date > NOW()
`

func (q *Queries) GetOriginalURL(ctx context.Context, customAlias string) (string, error) {
	row := q.db.QueryRowContext(ctx, getOriginalURL, customAlias)
	var original_url string
	err := row.Scan(&original_url)
	return original_url, err
}
