-- name: GetAuthorById :one
SELECT
  *
FROM
  author
WHERE
  id = $1;
