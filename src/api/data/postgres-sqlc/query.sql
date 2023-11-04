-- name: GetUrlByHash :one
SELECT * FROM urls
WHERE hash = $1 LIMIT 1;

-- name: GetUrlByID :one
SELECT * FROM urls
WHERE id = $1 LIMIT 1;

-- name: CreateUrl :one
INSERT into urls(originalurl, hash, created_at, owner) VALUES($1, $2, $3, $4) RETURNING *;