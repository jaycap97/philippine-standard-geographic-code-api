package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetBarangayList() ([]domain.Barangay, error) {
	var bs []domain.Barangay
	res := db.gorm.Find(&bs)
	if res.Error != nil {
		log.Println("Repository: failed to GET barangay list!")
		return bs, res.Error
	}
	return bs, nil
}

func (db *Database) GetBarangay(psgc_id string) (domain.Barangay, error) {
	var b domain.Barangay
	res := db.gorm.First(&b, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET barangay %s!", psgc_id))
		return b, res.Error
	}
	return b, nil
}

func (db *Database) CreateBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res := db.gorm.Create(&b)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE barangay %s!", b.PsgcID))
		return b, res.Error
	}
	return b, nil
}

func (db *Database) UpdateBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res := db.gorm.Save(&b)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE barangay %s!", b.PsgcID))
		return b, res.Error
	}
	return b, nil
}

func (db *Database) DeleteBarangay(b *domain.Barangay) (*domain.Barangay, error) {
	res := db.gorm.Delete(&b)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE barangay %s!", b.PsgcID))
		return b, res.Error
	}
	return b, nil
}
