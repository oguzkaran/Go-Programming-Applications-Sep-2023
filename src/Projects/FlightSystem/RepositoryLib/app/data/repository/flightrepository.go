package repository

import (
	"SampleTimeServiceApp/app/data/repository/entity"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type FlightRepository struct {
	db *gorm.DB
}

func NewFlightRepository(db *gorm.DB) *FlightRepository {
	return &FlightRepository{db}
}

func (fr *FlightRepository) Save(f *entity.Flight) {
	fr.db.Create(f)
}
