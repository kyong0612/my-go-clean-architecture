package repository

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/my-go-clean-architecture/domain"
	"github.com/kyong0612/my-go-clean-architecture/internal/repository/postgres"
)

func articleToModel(article postgres.Article, author *postgres.Author) *domain.Article {
	result := domain.Article{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}

	if author != nil {
		result.Author = authorToModel(*author)
	}

	return &result
}

func (r *Repository) FetchArticles(ctx context.Context, cursor string, num int32) ([]*domain.Article, string, error) {
	sqlCursor, err := DecodeCursor(cursor)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to decode cursor")
	}

	articles, err := r.db.ListArticles(ctx, postgres.ListArticlesParams{
		Cursor: sqlCursor,
		Limit:  int64(num),
	})
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to fetch articles")
	}

	result := make([]*domain.Article, 0, len(articles))
	for _, article := range articles {
		result = append(result, articleToModel(article, nil))
	}

	return result, EncodeCursor(result[len(result)-1].CreatedAt), nil
}

func (r *Repository) GetArticleByID(ctx context.Context, id int32) (*domain.Article, error) {
	article, err := r.db.GetArticleByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get article by id")
	}

	return articleToModel(article, nil), nil
}

func (r *Repository) GetArticleByTitle(ctx context.Context, title string) (*domain.Article, error) {
	article, err := r.db.GetArticleByTitle(ctx, title)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get article by title")
	}

	return articleToModel(article, nil), nil
}

func (r *Repository) UpdateArticle(ctx context.Context, ar domain.Article) error {
	if err := r.db.UpdateArticle(ctx, postgres.UpdateArticleParams{
		ID:      ar.ID,
		Title:   &ar.Title,
		Content: &ar.Content,
	}); err != nil {
		return errors.Wrap(err, "failed to update article")
	}

	return nil
}

func (r *Repository) StoreArticle(ctx context.Context, a domain.Article) error {
	param := postgres.CreateArticleParams{
		Title:   a.Title,
		Content: a.Content,
	}

	if a.Author != nil {
		param.AuthorID = &a.Author.ID
	}

	if err := r.db.CreateArticle(ctx, param); err != nil {
		return errors.Wrap(err, "failed to store article")
	}

	return nil
}

func (r *Repository) DeleteArticle(ctx context.Context, id int32) error {
	if err := r.db.DeleteArticle(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete article")
	}

	return nil
}
