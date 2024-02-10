package entity

import "gorm.io/gorm"

type PlaceInfo struct {
	gorm.Model
	Code               int    `gorm:"unique"`
	PlaceName          string `gorm:"unique"`
	Country            string
	East               float64
	North              float64
	South              float64
	West               float64
	WeatherInformation []WeatherInfo `gorm:"foreignKey:PlaceInfoId"`
}
