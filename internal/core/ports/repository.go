package ports

import "github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"

type Repository interface {
	GetRegionList() ([]domain.Region, error)
	GetRegion(psgc_id string) (domain.Region, error)
	CreateRegion(*domain.Region) (*domain.Region, error)
	UpdateRegion(*domain.Region) (*domain.Region, error)
	DeleteRegion(*domain.Region) (*domain.Region, error)

	GetProvinceList() ([]domain.Province, error)
	GetProvince(psgc_id string) (domain.Province, error)
	CreateProvince(*domain.Province) (*domain.Province, error)
	UpdateProvince(*domain.Province) (*domain.Province, error)
	DeleteProvince(*domain.Province) (*domain.Province, error)

	GetCityList() ([]domain.City, error)
	GetCity(psgc_id string) (domain.City, error)
	CreateCity(*domain.City) (*domain.City, error)
	UpdateCity(*domain.City) (*domain.City, error)
	DeleteCity(*domain.City) (*domain.City, error)

	GetMunicipalityList() ([]domain.Municipality, error)
	GetMunicipality(psgc_id string) (domain.Municipality, error)
	CreateMunicipality(*domain.Municipality) (*domain.Municipality, error)
	UpdateMunicipality(*domain.Municipality) (*domain.Municipality, error)
	DeleteMunicipality(*domain.Municipality) (*domain.Municipality, error)

	GetSubMunicipalityList() ([]domain.SubMunicipality, error)
	GetSubMunicipality(psgc_id string) (domain.SubMunicipality, error)
	CreateSubMunicipality(*domain.SubMunicipality) (*domain.SubMunicipality, error)
	UpdateSubMunicipality(*domain.SubMunicipality) (*domain.SubMunicipality, error)
	DeleteSubMunicipality(*domain.SubMunicipality) (*domain.SubMunicipality, error)

	GetBarangayList() ([]domain.Barangay, error)
	GetBarangay(psgc_id string) (domain.Barangay, error)
	CreateBarangay(*domain.Barangay) (*domain.Barangay, error)
	UpdateBarangay(*domain.Barangay) (*domain.Barangay, error)
	DeleteBarangay(*domain.Barangay) (*domain.Barangay, error)
}
