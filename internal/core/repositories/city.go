package repositories

import (
	"fmt"
	"log"

	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
)

func (db *Database) GetCityList() ([]domain.City, error) {
	var cs []domain.City
	res := db.gorm.Find(&cs)
	if res.Error != nil {
		log.Println("Repository: failed to GET city list!")
		return cs, res.Error
	}
	return cs, nil
}

func (db *Database) GetCity(psgc_id string) (domain.City, error) {
	var c domain.City
	res := db.gorm.First(&c, "psgc_id", psgc_id)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to GET city %s!", psgc_id))
		return c, res.Error
	}
	return c, nil
}

func (db *Database) CreateCity(c *domain.City) (*domain.City, error) {
	res := db.gorm.Create(&c)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to CREATE city %s!", c.PsgcID))
		return c, res.Error
	}
	return c, nil
}

func (db *Database) UpdateCity(c *domain.City) (*domain.City, error) {
	res := db.gorm.Save(&c)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to UPDATE city %s!", c.PsgcID))
		return c, res.Error
	}
	return c, nil
}

func (db *Database) DeleteCity(c *domain.City) (*domain.City, error) {
	res := db.gorm.Delete(&c)
	if res.Error != nil {
		log.Println(fmt.Printf("Repository: failed to DELETE city %s!", c.PsgcID))
		return c, res.Error
	}
	return c, nil
}
