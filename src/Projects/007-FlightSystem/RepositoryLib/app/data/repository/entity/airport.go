package entity

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	City               string
	DepartureFlights   []Flight `gorm:"foreignKey:DepartureAirportID"`
	DestinationFlights []Flight `gorm:"foreignKey:DestinationAirportID"`
}
