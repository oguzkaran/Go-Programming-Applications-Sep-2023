package repository

import (
	"WeatherInfoScheduler/data/entity"
	"gorm.io/gorm"
)

type WeatherRepository struct {
	db *gorm.DB
}

func NewWeatherRepository(db *gorm.DB) *WeatherRepository {
	return &WeatherRepository{db}
}

func (wr *WeatherRepository) Save(weather *entity.Weather) error {
	result := wr.db.Create(weather)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (wr *WeatherRepository) Delete(weather *entity.Weather) error {
	result := wr.db.Delete(weather)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
