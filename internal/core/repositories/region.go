package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetRegionList() ([]domain.Region, error) {
	var rs []domain.Region
	res := db.gorm.Find(&rs)
	if res.Error != nil {
		log.Println("Repository: failed to GET region list!")
		return rs, res.Error
	}
	return rs, nil
}

func (db *Database) GetRegion(psgc_id string) (domain.Region, error) {
	var r domain.Region
	res := db.gorm.First(&r, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET region %s!", psgc_id))
		return r, res.Error
	}
	return r, nil
}

func (db *Database) CreateRegion(r *domain.Region) (*domain.Region, error) {
	res := db.gorm.Create(&r)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE region %s!", r.PsgcID))
		return r, res.Error
	}
	return r, nil
}

func (db *Database) UpdateRegion(r *domain.Region) (*domain.Region, error) {
	res := db.gorm.Save(&r)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE region %s!", r.PsgcID))
		return r, res.Error
	}
	return r, nil
}

func (db *Database) DeleteRegion(r *domain.Region) (*domain.Region, error) {
	res := db.gorm.Delete(&r)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE region %s!", r.PsgcID))
		return r, res.Error
	}
	return r, nil
}
