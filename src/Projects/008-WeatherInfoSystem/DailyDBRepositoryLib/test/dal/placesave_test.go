package dal

import (
	"WeatherInfoScheduler/data/dal"
	"WeatherInfoScheduler/data/entity"
	"testing"
)

func TestSavePlace(t *testing.T) {
	helper, err := dal.NewWeatherInfoHelper()

	if err != nil {
		t.Errorf("DAL object creation problem occurred:%s\n", err.Error())
		return
	}
	placeInfo := &entity.Place{Code: 67, PlaceName: "Zonguldak", Country: "TR", North: 23.4, South: 34.5, East: 56.7, West: 45.7}

	err = helper.SavePlace(placeInfo)

	if err != nil {
		t.Errorf("Save problem occurred:%s\n", err.Error())
		return
	}
}
