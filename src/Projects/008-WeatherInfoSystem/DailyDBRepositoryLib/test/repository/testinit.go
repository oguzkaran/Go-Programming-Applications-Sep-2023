package repository

import (
	"WeatherInfoScheduler/data/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	const URL = "postgres://postgres:csd1993@localhost:5432/gs23_weatherappdb_dev"
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Place{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Weather{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
