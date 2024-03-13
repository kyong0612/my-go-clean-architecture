package repository

import "github.com/kyong0612/my-go-clean-architecture/internal/repository/postgres"

type Repository struct {
	db *postgres.Queries
}

func New(db postgres.DBTX) *Repository {
	return &Repository{
		db: postgres.New(db),
	}
}
