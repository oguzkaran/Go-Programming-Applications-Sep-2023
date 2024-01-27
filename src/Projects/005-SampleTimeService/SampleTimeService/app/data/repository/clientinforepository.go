package repository

import (
	"SampleTimeServiceApp/app/data/repository/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type ClientInfoRepository struct {
	db *gorm.DB
}

func NewClientInfoRepository() (*ClientInfoRepository, error) {
	db, err := gorm.Open("postgres", "host=csd-postgresql-db.cmxkkfycsomh.us-east-1.rds.amazonaws.com port=5432 user=postgres password=csystem1993 sslmode=disable dbname=gs23_timeservicedb")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.ClientInfo{})

	return &ClientInfoRepository{db}, nil
}

func (cr *ClientInfoRepository) Close() error {
	return cr.db.Close()
}

func (cr *ClientInfoRepository) Save(ci *entity.ClientInfo) {
	cr.db.Create(ci)
}
