package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetSubMunicipalityList() ([]domain.SubMunicipality, error) {
	res, err := r.repo.GetSubMunicipalityList()
	return res, err
}

func (r *Repository) GetSubMunicipality(psgc_id string) (domain.SubMunicipality, error) {
	res, err := r.repo.GetSubMunicipality(psgc_id)
	return res, err
}

func (r *Repository) CreateSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res, err := r.repo.CreateSubMunicipality(sm)
	return res, err
}

func (r *Repository) UpdateSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res, err := r.repo.UpdateSubMunicipality(sm)
	return res, err
}

func (r *Repository) DeleteSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res, err := r.repo.DeleteSubMunicipality(sm)
	return res, err
}
