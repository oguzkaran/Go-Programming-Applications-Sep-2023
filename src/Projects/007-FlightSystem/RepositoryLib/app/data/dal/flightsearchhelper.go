package dal

import (
	"SampleTimeServiceApp/app/data/repository"
	"SampleTimeServiceApp/app/data/repository/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FlightSearchHelper struct {
	ar *repository.AirPortRepository
	fr *repository.FlightRepository
}

func NewFlightSearchHelper() (*FlightSearchHelper, error) {
	const URL = "postgres://postgres:csystem1993@csd-postgresql-db.cmxkkfycsomh.us-east-1.rds.amazonaws.com:5432/gs23_flightsearchdb"
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Airport{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entity.Flight{})

	if err != nil {
		return nil, err
	}

	return &FlightSearchHelper{repository.NewAirPortRepository(db), repository.NewFlightRepository(db)}, nil
}

func (h *FlightSearchHelper) CountAirPort() int64 {
	return h.ar.Count()
}

func (h *FlightSearchHelper) SaveAirPort(a *entity.Airport) {
	h.ar.Save(a)
}

func (h *FlightSearchHelper) SaveFlight(f *entity.Flight) {
	h.fr.Save(f)
}
