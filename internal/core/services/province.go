package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetProvinceList() ([]domain.Province, error) {
	res, err := r.repo.GetProvinceList()
	return res, err
}

func (r *Repository) GetProvince(psgc_id string) (domain.Province, error) {
	res, err := r.repo.GetProvince(psgc_id)
	return res, err
}

func (r *Repository) CreateProvince(p *domain.Province) (*domain.Province, error) {
	res, err := r.repo.CreateProvince(p)
	return res, err
}

func (r *Repository) UpdateProvince(p *domain.Province) (*domain.Province, error) {
	res, err := r.repo.UpdateProvince(p)
	return res, err
}

func (r *Repository) DeleteProvince(p *domain.Province) (*domain.Province, error) {
	res, err := r.repo.DeleteProvince(p)
	return res, err
}
