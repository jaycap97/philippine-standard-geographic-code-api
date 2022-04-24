package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetCityList() ([]domain.City, error) {
	res, err := r.repo.GetCityList()
	return res, err
}

func (r *Repository) GetCity(psgc_id string) (domain.City, error) {
	res, err := r.repo.GetCity(psgc_id)
	return res, err
}

func (r *Repository) CreateCity(c *domain.City) (*domain.City, error) {
	res, err := r.repo.CreateCity(c)
	return res, err
}

func (r *Repository) UpdateCity(c *domain.City) (*domain.City, error) {
	res, err := r.repo.UpdateCity(c)
	return res, err
}

func (r *Repository) DeleteCity(c *domain.City) (*domain.City, error) {
	res, err := r.repo.DeleteCity(c)
	return res, err
}
