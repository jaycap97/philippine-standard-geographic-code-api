package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetRegionList() ([]domain.Region, error) {
	res, err := r.repo.GetRegionList()
	return res, err
}

func (r *Repository) GetRegion(psgc_id string) (domain.Region, error) {
	res, err := r.repo.GetRegion(psgc_id)
	return res, err
}

func (r *Repository) CreateRegion(rg *domain.Region) (*domain.Region, error) {
	res, err := r.repo.CreateRegion(rg)
	return res, err
}

func (r *Repository) UpdateRegion(rg *domain.Region) (*domain.Region, error) {
	res, err := r.repo.UpdateRegion(rg)
	return res, err
}

func (r *Repository) DeleteRegion(rg *domain.Region) (*domain.Region, error) {
	res, err := r.repo.DeleteRegion(rg)
	return res, err
}
