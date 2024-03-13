package article

import (
	"context"

	"github.com/kyong0612/my-go-clean-architecture/domain"
)

// AuthorRepository represent the author's repository contract
//
//go:generate mockery --name AuthorRepository
type AuthorRepository interface {
	GetAuthorByID(ctx context.Context, id int32) (domain.Author, error)
}
