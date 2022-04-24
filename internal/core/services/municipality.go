package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetMunicipalityList() ([]domain.Municipality, error) {
	res, err := r.repo.GetMunicipalityList()
	return res, err
}

func (r *Repository) GetMunicipality(psgc_id string) (domain.Municipality, error) {
	res, err := r.repo.GetMunicipality(psgc_id)
	return res, err
}

func (r *Repository) CreateMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res, err := r.repo.CreateMunicipality(m)
	return res, err
}

func (r *Repository) UpdateMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res, err := r.repo.UpdateMunicipality(m)
	return res, err
}

func (r *Repository) DeleteMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res, err := r.repo.DeleteMunicipality(m)
	return res, err
}
