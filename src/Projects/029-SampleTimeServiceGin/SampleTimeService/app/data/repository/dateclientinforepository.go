package repository

import (
	"SampleTimeServiceApp/app/data/repository/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type DateClientInfoRepository struct {
	db *gorm.DB
}

func NewDateClientInfoRepository(db *gorm.DB) *DateClientInfoRepository {
	return &DateClientInfoRepository{db}
}

func (cr *DateClientInfoRepository) Save(ci *entity.DateClientInfo) {
	cr.db.Create(ci)
}
