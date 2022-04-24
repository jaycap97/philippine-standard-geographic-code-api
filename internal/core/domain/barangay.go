package domain

type Barangay struct {
	PsgcID                string `json:"psgc_id" gorm:"primary_key;not_null;type:char(10)"`
	Name                  string `json:"name" gorm:"type:varchar(100)"`
	CorrespondenceCode    string `json:"correspondence_code" gorm:"type:char(9);default:null"`
	OldName               string `json:"old_name" gorm:"type:varchar(100);default:null"`
	MunicipalityPsgcID    string `json:"municipality_psgc_id"`
	CityPsgcID            string `json:"city_psgc_id"`
	SubMunicipalityPsgcID string `json:"sub_municipality_psgc_id"`
}
