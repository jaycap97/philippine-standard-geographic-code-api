package domain

type Province struct {
	PsgcID             string            `json:"psgc_id" gorm:"primary_key;not_null;type:char(10)"`
	Name               string            `json:"name" gorm:"type:varchar(100)"`
	CorrespondenceCode string            `json:"correspondence_code" gorm:"type:char(9);default:null"`
	OldName            string            `json:"old_name" gorm:"type:varchar(100);default:null"`
	Cities             []City            `json:"cities,omitempty" gorm:"foreignKey:ProvincePsgcID;references:PsgcID"`
	Municipalities     []Municipality    `json:"municipalities,omitempty" gorm:"foreignKey:ProvincePsgcID;references:PsgcID"`
	SubMunicipalities  []SubMunicipality `json:"sub-municipalities,omitempty" gorm:"foreignKey:ProvincePsgcID;references:PsgcID"`
	RegionPsgcID       string            `json:"region_psgc_id"`
}
