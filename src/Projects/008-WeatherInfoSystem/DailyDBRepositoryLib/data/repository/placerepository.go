package repository

import (
	"WeatherInfoScheduler/data/entity"
	"errors"
	"gorm.io/gorm"
)

type PlaceRepository struct {
	db *gorm.DB
}

func NewPlaceRepository(db *gorm.DB) *PlaceRepository {
	return &PlaceRepository{db}
}

func (pr *PlaceRepository) Save(place *entity.Place) error {
	result := pr.db.Create(place)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (pr *PlaceRepository) FindByPlaceName(placeName string) (*entity.Place, error) {
	var pi entity.Place

	result := pr.db.Where("place_name = ?", placeName).First(&pi)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &pi, nil
}

func (pr *PlaceRepository) FindByCode(code int) (*entity.Place, error) {
	var pi entity.Place

	result := pr.db.Where("code = ?", code).First(&pi)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &pi, nil
}
