package models

import (
	"github.com/msajad79/pricing_engine/config"
	"gorm.io/gorm"
)

func MigrateModels() {
	config.DB.AutoMigrate(&Rule{}, &Route{}, &City{}, &Supplier{}, &Airline{}, &Agency{})
}

type Rule struct {
	ID          uint64     ``
	Routes      []Route    `json:"rutes"         gorm:"many2many:rule_cities"  binding:"is_valid_routes"`
	Airlines    []Airline  `json:"airline"       gorm:"many2many:rule_airlines" `
	Agencies    []Agency   `json:"agency"        gorm:"many2many:rule_agencies" `
	Suppliers   []Supplier `json:"supplier"      gorm:"many2many:rule_supplier" `
	AmountType  string     `json:"amountType"    gorm:"varchar(64)"             `
	AmountValue float64    `json:"amountValue"   gorm:"double"                  `
}

type Route struct {
	ID          uint64 ``
	Origin      string `json:"origin"      gorm:"varchar(3)" binding:"custombinding"`
	Destination string `json:"destination" gorm:"varchar(3)"`
}

type City struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(3)"`
}

type Supplier struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(64)"`
}

type Airline struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(64)"`
}

type Agency struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(64)"`
}
