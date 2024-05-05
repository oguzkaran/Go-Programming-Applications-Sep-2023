package entity

import "github.com/jinzhu/gorm"

type PlaceInfo struct {
	gorm.Model
	Name      string
	Latitude  float64
	Longitude float64
}
