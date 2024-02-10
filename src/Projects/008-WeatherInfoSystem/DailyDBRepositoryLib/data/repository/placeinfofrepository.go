package repository

import (
	"WeatherInfoScheduler/data/entity"
	"gorm.io/gorm"
)

type PlaceInfoRepository struct {
	db *gorm.DB
}

func NewPlaceInfoRepository(db *gorm.DB) *PlaceInfoRepository {
	return &PlaceInfoRepository{db}
}

func (pr *PlaceInfoRepository) Save(placeInfo *entity.PlaceInfo) error {
	result := pr.db.Create(placeInfo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *PlaceInfoRepository) FindByPlaceName(placeName string) (*entity.PlaceInfo, error) {
	var pi entity.PlaceInfo

	result := pr.db.First(&pi, placeName)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pi, nil
}
