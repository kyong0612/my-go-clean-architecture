-- name: GetAuthorByID :one
SELECT
  *
FROM
  author
WHERE
  id = $1;
