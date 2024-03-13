-- name: ListAuthors :many
SELECT
  *
FROM
  article
WHERE
  created_at > $1
ORDER BY
  created_at
LIMIT
  $2;

-- name: GetArticleByID :one
SELECT
  *
FROM
  article
WHERE
  id = $1;

-- name: GetArticleByTitle :one
SELECT
  *
FROM
  article
WHERE
  title LIKE $1
LIMIT
  1;

-- name: CreateArticle :exec
INSERT INTO article(title, content, author_id) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateArticles :batchexec
INSERT INTO article(title, content, author_id) VALUES ($1, $2, $3);

-- name: UpdateArticle :exec
UPDATE article
SET
  title = coalesce(sqlc.narg('title'), title),
  content = coalesce(sqlc.narg('content'), content)
WHERE
  id = sqlc.arg('id');

-- name: DeleteArticle :exec
DELETE FROM article
WHERE
  id = $1;
