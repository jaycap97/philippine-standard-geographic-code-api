package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetMunicipalityList() ([]domain.Municipality, error) {
	var ms []domain.Municipality
	res := db.gorm.Find(&ms)
	if res.Error != nil {
		log.Println("Repository: failed to GET municipality list!")
		return ms, res.Error
	}
	return ms, nil
}

func (db *Database) GetMunicipality(psgc_id string) (domain.Municipality, error) {
	var m domain.Municipality
	res := db.gorm.First(&m, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET municipality %s!", psgc_id))
		return m, res.Error
	}
	return m, nil
}

func (db *Database) CreateMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res := db.gorm.Create(&m)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE municipality %s!", m.PsgcID))
		return m, res.Error
	}
	return m, nil
}

func (db *Database) UpdateMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res := db.gorm.Save(&m)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE municipality %s!", m.PsgcID))
		return m, res.Error
	}
	return m, nil
}

func (db *Database) DeleteMunicipality(m *domain.Municipality) (*domain.Municipality, error) {
	res := db.gorm.Delete(&m)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE municipality %s!", m.PsgcID))
		return m, res.Error
	}
	return m, nil
}
