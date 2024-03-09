package dal

import (
	"WeatherInfoScheduler/data/entity"
	"WeatherInfoScheduler/data/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WeatherInfoHelper struct {
	db *gorm.DB
	wr *repository.WeatherRepository
	pr *repository.PlaceRepository
}

func initDb() (*gorm.DB, error) {
	const URL = "postgres://postgres:csystem1993@csd-postgresql-db.cmxkkfycsomh.us-east-1.rds.amazonaws.com:5432/gs23_weatherappdb"
	db, err := gorm.Open(postgres.Open(URL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Place{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Weather{})
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

	return &WeatherInfoHelper{db, repository.NewWeatherRepository(db), repository.NewPlaceRepository(db)}, nil
}

func (wh *WeatherInfoHelper) DeleteWeather(weatherInfo *entity.Weather) error {
	return wh.wr.Delete(weatherInfo)
}

func (wh *WeatherInfoHelper) FindPlaceByPlaceName(placeName string) (*entity.Place, error) {
	return wh.pr.FindByPlaceName(placeName)
}

func (wh *WeatherInfoHelper) FindPlaceByCode(code int) (*entity.Place, error) {
	return wh.pr.FindByCode(code)
}

func (wh *WeatherInfoHelper) SavePlace(placeInfo *entity.Place) error {
	return wh.pr.Save(placeInfo)
}

func (wh *WeatherInfoHelper) SaveWeather(weatherInfo *entity.Weather) error {
	return wh.wr.Save(weatherInfo)
}

func (wh *WeatherInfoHelper) ExecNonQuery(sqlCmd string) error {
	return wh.db.Exec(sqlCmd).Error
}

//...
