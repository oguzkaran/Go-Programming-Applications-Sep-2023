package repository

import (
	"SampleTimeServiceApp/app/data/repository/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type TimeClientInfoRepository struct {
	db *gorm.DB
}

func NewTimeClientInfoRepository(db *gorm.DB) *TimeClientInfoRepository {
	return &TimeClientInfoRepository{db}
}

func (cr *TimeClientInfoRepository) Close() error {
	return cr.db.Close()
}

func (cr *TimeClientInfoRepository) Save(ci *entity.TimeClientInfo) {
	cr.db.Create(ci)
}
