package repository

import (
	"PlaceInfoInserter/app/data/entity"
	"github.com/jinzhu/gorm"
)

type PlaceInfoRepository struct {
	db *gorm.DB
}

func (pr *PlaceInfoRepository) Save(pi *entity.PlaceInfo) *entity.PlaceInfo {
	pr.db.Create(pi)

	return pi
}
