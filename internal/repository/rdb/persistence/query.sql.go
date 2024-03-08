// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package persistence

import (
	"context"
)

const getArticle = `-- name: GetArticle :one
SELECT id, title, content, author_id, updated_at, created_at FROM article
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArticle(ctx context.Context, id int32) (Article, error) {
	row := q.db.QueryRow(ctx, getArticle, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
