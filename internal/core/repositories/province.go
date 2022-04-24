package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetProvinceList() ([]domain.Province, error) {
	var ps []domain.Province
	res := db.gorm.Find(&ps)
	if res.Error != nil {
		log.Println("Repository: failed to GET province list!")
		return ps, res.Error
	}
	return ps, nil
}

func (db *Database) GetProvince(psgc_id string) (domain.Province, error) {
	var p domain.Province
	res := db.gorm.First(&p, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET province %s!", psgc_id))
		return p, res.Error
	}
	return p, nil
}

func (db *Database) CreateProvince(p *domain.Province) (*domain.Province, error) {
	res := db.gorm.Create(&p)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE province %s!", p.PsgcID))
		return p, res.Error
	}
	return p, nil
}

func (db *Database) UpdateProvince(p *domain.Province) (*domain.Province, error) {
	res := db.gorm.Save(&p)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE province %s!", p.PsgcID))
		return p, res.Error
	}
	return p, nil
}

func (db *Database) DeleteProvince(p *domain.Province) (*domain.Province, error) {
	res := db.gorm.Delete(&p)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE province %s!", p.PsgcID))
		return p, res.Error
	}
	return p, nil
}
