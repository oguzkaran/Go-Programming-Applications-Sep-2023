package repository

import (
	"WeatherInfoScheduler/data/entity"
	"gorm.io/gorm"
)

type WeatherInfoRepository struct {
	db *gorm.DB
}

func NewWeatherInfoRepository(db *gorm.DB) *WeatherInfoRepository {
	return &WeatherInfoRepository{db}
}

func (wr *WeatherInfoRepository) Save(weatherInfo *entity.WeatherInfo) error {
	result := wr.db.Create(weatherInfo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (wr *WeatherInfoRepository) Delete(weatherInfo *entity.WeatherInfo) error {
	result := wr.db.Delete(weatherInfo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
