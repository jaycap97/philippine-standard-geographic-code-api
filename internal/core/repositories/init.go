package repositories

import (
	"github.com/jaycap97/philippine-standard-geographic-code-api/internal/core/domain"
	"gorm.io/gorm"
)

type Database struct {
	gorm *gorm.DB
}

func Init(gorm *gorm.DB) *Database {
	gorm.AutoMigrate(
		&domain.Region{},
		&domain.Province{},
		&domain.City{},
		&domain.Municipality{},
		&domain.SubMunicipality{},
		&domain.Barangay{},
	)
	return &Database{gorm}
}
