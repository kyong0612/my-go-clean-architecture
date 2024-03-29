package repository

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/kyong0612/my-go-clean-architecture/domain"
	"github.com/kyong0612/my-go-clean-architecture/internal/repository/postgres"
)

func authorToModel(author postgres.Author) *domain.Author {
	return &domain.Author{
		ID:        author.ID,
		Name:      author.Name,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}

func (r *Repository) GetAuthorByID(ctx context.Context, id int32) (*domain.Author, error) {
	author, err := r.db.GetAuthorByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get author by id")
	}

	return authorToModel(author), nil
}
