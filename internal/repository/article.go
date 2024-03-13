package repository

import (
	"context"

	"github.com/kyong0612/my-go-clean-architecture/domain"
	"github.com/kyong0612/my-go-clean-architecture/internal/repository/postgres"
)

func articleToModel(article postgres.Article, author postgres.Author) domain.Article {
	return domain.Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Author:    authorToModel(author),
		CreatedAt: article.CreatedAt.Time,
		UpdatedAt: article.UpdatedAt.Time,
	}
}

func (r *Repository) FetchArticles(ctx context.Context, cursor string, num int32) (res []domain.Article, nextCursor string, err error) {
	// TODO: implement the code
	return nil, "", nil
}

func (r *Repository) GetArticleByID(ctx context.Context, id int32) (domain.Article, error) {
	// TODO: implement the code
	return domain.Article{}, nil
}

func (r *Repository) GetArticleByTitle(ctx context.Context, title string) (domain.Article, error) {
	// TODO: implement the code
	return domain.Article{}, nil
}

func (r *Repository) UpdateArticle(ctx context.Context, ar *domain.Article) error {
	// TODO: implement the code
	return nil
}

func (r *Repository) StoreArticle(ctx context.Context, a *domain.Article) error {
	// TODO: implement the code
	return nil
}

func (r *Repository) DeleteArticle(ctx context.Context, id int32) error {
	// TODO: implement the code
	return nil
}
