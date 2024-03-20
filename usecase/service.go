package usecase

import (
	"github.com/kyong0612/my-go-clean-architecture/internal/repository"
)

var _ Repository = (*repository.Repository)(nil)

type Repository interface {
	AuthorRepository
	ArticleRepository
}

type Service struct {
	repo Repository
}

// NewService will create a new article service object.
func NewService(repo Repository) *Service {
	return &Service{
		repo,
	}
}
