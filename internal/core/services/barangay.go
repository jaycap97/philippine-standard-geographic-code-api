package services

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (r *Repository) GetBarangayList() ([]domain.Barangay, error) {
	res, err := r.repo.GetBarangayList()
	return res, err
}

func (r *Repository) GetBarangay(psgc_id string) (domain.Barangay, error) {
	res, err := r.repo.GetBarangay(psgc_id)
	return res, err
}

func (r *Repository) CreateBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res, err := r.repo.CreateBarangay(b)
	return res, err
}

func (r *Repository) UpdateBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res, err := r.repo.UpdateBarangay(b)
	return res, err
}

func (r *Repository) DeleteBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res, err := r.repo.DeleteBarangay(b)
	return res, err
}
