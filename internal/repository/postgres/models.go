// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgres

import (
	"time"
)

type Article struct {
	ID        int32
	Title     string
	Content   string
	AuthorID  *int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

type ArticleCategory struct {
	ID         int32
	ArticleID  int32
	CategoryID int32
}

type Author struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Category struct {
	ID        int32
	Name      string
	Tag       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
