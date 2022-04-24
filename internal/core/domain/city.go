package domain

type City struct {
	PsgcID             string     `json:"psgc_id" gorm:"primary_key;not_null;type:char(10)"`
	Name               string     `json:"name" gorm:"type:varchar(100)"`
	CorrespondenceCode string     `json:"correspondence_code" gorm:"type:char(9);default:null"`
	OldName            string     `json:"old_name" gorm:"type:varchar(100);default:null"`
	Barangays          []Barangay `json:"barangays,omitempty" gorm:"foreignKey:CityPsgcID;references:PsgcID"`
	ProvincePsgcID     string     `json:"province_psgc_id"`
}
