package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetSubMunicipalityList() ([]domain.SubMunicipality, error) {
	var sms []domain.SubMunicipality
	res := db.gorm.Find(&sms)
	if res.Error != nil {
		log.Println("Repository: failed to GET sub-municipality list!")
		return sms, res.Error
	}
	return sms, nil
}

func (db *Database) GetSubMunicipality(psgc_id string) (domain.SubMunicipality, error) {
	var sm domain.SubMunicipality
	res := db.gorm.First(&sm, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET sub-municipality %s!", psgc_id))
		return sm, res.Error
	}
	return sm, nil
}

func (db *Database) CreateSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res := db.gorm.Create(&sm)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE sub-municipality %s!", sm.PsgcID))
		return sm, res.Error
	}
	return sm, nil
}

func (db *Database) UpdateSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res := db.gorm.Save(&sm)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE sub-municipality %s!", sm.PsgcID))
		return sm, res.Error
	}
	return sm, nil
}

func (db *Database) DeleteSubMunicipality(sm *domain.SubMunicipality) (*domain.SubMunicipality, error) {
	res := db.gorm.Delete(&sm)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE sub-municipality %s!", sm.PsgcID))
		return sm, res.Error
	}
	return sm, nil
}
