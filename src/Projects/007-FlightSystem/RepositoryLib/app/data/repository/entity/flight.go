package entity

import (
	"gorm.io/gorm"
	"time"
)

type Flight struct {
	gorm.Model
	DepartureAirportID   int
	DestinationAirportID int
	DepartureTime        time.Time
	DestinationTime      time.Time
	Price                float64
}
