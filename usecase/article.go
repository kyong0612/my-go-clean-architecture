package usecase

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/my-go-clean-architecture/domain"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// ArticleRepository represent the article's repository contract
//
//go:generate mockery --name ArticleRepository
type ArticleRepository interface {
	FetchArticles(ctx context.Context, cursor string, num int32) (res []*domain.Article, nextCursor string, err error)
	GetArticleByID(ctx context.Context, id int32) (*domain.Article, error)
	GetArticleByTitle(ctx context.Context, title string) (*domain.Article, error)
	UpdateArticle(ctx context.Context, ar domain.Article) error
	StoreArticle(ctx context.Context, a domain.Article) error
	DeleteArticle(ctx context.Context, id int32) error
}

/*
* In this function below, I'm using errgroup with the pipeline pattern
* Look how this works in this package explanation
* in godoc: https://godoc.org/golang.org/x/sync/errgroup#ex-Group--Pipeline
 */
func (a *Service) fillAuthorDetails(ctx context.Context, data []*domain.Article) ([]*domain.Article, error) {
	g, ctx := errgroup.WithContext(ctx)
	// Get the author's id
	mapAuthors := map[int32]domain.Author{}

	for _, article := range data {
		mapAuthors[article.Author.ID] = domain.Author{}
	}
	// Using goroutine to fetch the author's detail
	chanAuthor := make(chan *domain.Author)
	for authorID := range mapAuthors {
		authorID := authorID
		g.Go(func() error {
			res, err := a.repo.GetAuthorByID(ctx, authorID)
			if err != nil {
				return err
			}
			chanAuthor <- res
			return nil
		})
	}

	go func() {
		defer close(chanAuthor)
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
	}()

	for author := range chanAuthor {
		if author != nil {
			mapAuthors[author.ID] = *author
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	// merge the author's data
	for index, item := range data {
		if a, ok := mapAuthors[item.Author.ID]; ok {
			data[index].Author = &a
		}
	}
	return data, nil
}

func (a *Service) FetchArticles(ctx context.Context, cursor string, num int32) (res []*domain.Article, nextCursor string, err error) {
	res, nextCursor, err = a.repo.FetchArticles(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	res, err = a.fillAuthorDetails(ctx, res)
	if err != nil {
		nextCursor = ""
	}

	return res, nextCursor, err
}

func (a *Service) GetArticleByID(ctx context.Context, id int32) (*domain.Article, error) {
	res, err := a.repo.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.repo.GetAuthorByID(ctx, res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = resAuthor

	return res, err
}

func (a *Service) UpdateArticle(ctx context.Context, ar *domain.Article) (err error) {
	if ar == nil {
		return nil
	}

	ar.UpdatedAt = time.Now()
	return a.repo.UpdateArticle(ctx, *ar)
}

func (a *Service) GetArticleByTitle(ctx context.Context, title string) (*domain.Article, error) {
	res, err := a.repo.GetArticleByTitle(ctx, title)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get article by title")
	}

	resAuthor, err := a.repo.GetAuthorByID(ctx, res.Author.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get author by id")
	}

	res.Author = resAuthor
	return res, err
}

func (a *Service) StoreArticle(ctx context.Context, m *domain.Article) error {
	if m == nil {
		return nil
	}

	return a.repo.StoreArticle(ctx, *m)
}

func (a *Service) DeleteArticle(ctx context.Context, id int32) error {
	return a.repo.DeleteArticle(ctx, id)
}
