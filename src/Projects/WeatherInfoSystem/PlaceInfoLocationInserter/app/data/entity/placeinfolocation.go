package entity

import "github.com/jinzhu/gorm"

type PlaceInfoLocation struct {
	gorm.Model
	Name      string
	Latitude  float64
	Longitude float64
}
