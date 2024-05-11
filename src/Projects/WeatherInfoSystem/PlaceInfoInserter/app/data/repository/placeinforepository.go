package repository

import (
	"PlaceInfoInserter/app/data/entity"
	"github.com/jinzhu/gorm"
)

type PlaceInfoRepository struct {
	db *gorm.DB
}

const dialect = "postgres"
const url = "host=localhost port=5432 user=postgres password=csd1993 sslmode=disable dbname=gs23_weatherinfosystemdb_dev"

func initDB() *PlaceInfoRepository {
	db, err := gorm.Open(dialect, url)
	if err != nil {
		return nil
	}
	db.AutoMigrate(&entity.PlaceInfo{})

	return NewPlaceInfoRepositoryDB(db)
}

func NewPlaceInfoRepository() *PlaceInfoRepository {
	//...
	return initDB()
}

func NewPlaceInfoRepositoryDB(db *gorm.DB) *PlaceInfoRepository {
	return &PlaceInfoRepository{db}
}

func NewPlaceInfoRepositoryDbInfo(host string, port int, username, password string) *PlaceInfoRepository {
	//...
	return &PlaceInfoRepository{}
}

func (pr *PlaceInfoRepository) Save(pi *entity.PlaceInfo) *entity.PlaceInfo {
	pr.db.Create(pi)

	return pi
}
