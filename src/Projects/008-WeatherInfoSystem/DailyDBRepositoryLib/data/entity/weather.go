package entity

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	PlaceInfoId      int
	Lng              float64
	Observation      string
	Icao             string
	Clouds           string
	DewPoint         string
	Datetime         string
	Temperature      string
	Humidity         float64
	StationName      string
	WeatherCondition string
	WindDirection    float64
	WindSpeed        string
	Lat              float64
	CloudsCode       string
}
