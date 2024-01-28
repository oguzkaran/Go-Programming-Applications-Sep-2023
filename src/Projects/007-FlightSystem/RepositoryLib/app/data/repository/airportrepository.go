package repository

import (
	"SampleTimeServiceApp/app/data/repository/entity"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type AirPortRepository struct {
	db *gorm.DB
}

func NewAirPortRepository(db *gorm.DB) *AirPortRepository {
	return &AirPortRepository{db}
}

func (ar *AirPortRepository) Save(a *entity.Airport) {
	ar.db.Create(a)
}

func (ar *AirPortRepository) Count() int64 {
	var count int64

	ar.db.Count(&count)

	return count
}
