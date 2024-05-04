package dal

import (
	"SampleTimeServiceApp/app/data/repository"
	"SampleTimeServiceApp/app/data/repository/entity"
	"github.com/jinzhu/gorm"
)

type TimeServiceHelper struct {
	tcr *repository.TimeClientInfoRepository
	dcr *repository.DateClientInfoRepository
}

func NewTimeServiceHelper() (*TimeServiceHelper, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=csd1993 sslmode=disable dbname=gs23_timeservicedb")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.TimeClientInfo{})
	db.AutoMigrate(&entity.DateClientInfo{})

	return &TimeServiceHelper{repository.NewTimeClientInfoRepository(db), repository.NewDateClientInfoRepository(db)}, nil
}

func (tsh *TimeServiceHelper) SaveTimeClientInfo(ci *entity.TimeClientInfo) {
	tsh.tcr.Save(ci)
}

func (tsh *TimeServiceHelper) SaveDateClientInfo(ci *entity.DateClientInfo) {
	tsh.dcr.Save(ci)
}
