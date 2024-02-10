package dal

import (
	entity2 "WeatherInfoScheduler/data/entity"
	repository2 "WeatherInfoScheduler/data/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WeatherInfoHelper struct {
	wr *repository2.WeatherInfoRepository
	pr *repository2.PlaceInfoRepository
}

func initDb() (*gorm.DB, error) {
	const URL = "postgres://postgres:csystem1993@csd-postgresql-db.cmxkkfycsomh.us-east-1.rds.amazonaws.com:5432/gs23_weatherappdb"
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity2.PlaceInfo{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity2.WeatherInfo{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewWeatherInfoHelper() (*WeatherInfoHelper, error) {
	db, err := initDb()

	if err != nil {
		return nil, err
	}

	return &WeatherInfoHelper{repository2.NewWeatherInfoRepository(db), repository2.NewPlaceInfoRepository(db)}, nil
}

func (wh *WeatherInfoHelper) SavePlaceInfo(placeInfo *entity2.PlaceInfo) error {
	return wh.pr.Save(placeInfo)
}

func (wh *WeatherInfoHelper) SaveWeatherInfo(weatherInfo *entity2.WeatherInfo) error {
	return wh.wr.Save(weatherInfo)
}

func (wh *WeatherInfoHelper) DeleteWeatherInfo(weatherInfo *entity2.WeatherInfo) error {
	return wh.wr.Delete(weatherInfo)
}

//...
