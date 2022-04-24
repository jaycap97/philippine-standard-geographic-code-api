package domain

type Region struct {
	PsgcID             string     `json:"psgc_id" gorm:"primary_key;not_null;type:char(10)"`
	Name               string     `json:"name" gorm:"type:varchar(100)"`
	CorrespondenceCode string     `json:"correspondence_code" gorm:"type:char(9);default:null"`
	OldName            string     `json:"old_name" gorm:"type:varchar(100);default:null"`
	Provinces          []Province `json:"provinces,omitempty" gorm:"foreignKey:RegionPsgcID;references:PsgcID"`
}
