package services

import "github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/repositories"

type Repository struct {
	repo *repositories.Database
}

func Init(repo *repositories.Database) *Repository {
	return &Repository{repo}
}
