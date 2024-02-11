package dal

import (
	"WeatherInfoScheduler/data/dal"
	"testing"
)

func TestFindPlaceByPlaceName(t *testing.T) {
	helper, err := dal.NewWeatherInfoHelper()

	if err != nil {
		t.Errorf("DAL object creation problem occurred:%s\n", err.Error())
		return
	}

	if helper.ExecNonQuery("truncate table places restart identity cascade") != nil {
		t.Errorf("DB problem occurred:%s\n", err.Error())
		return
	}

	TestSavePlace(t)

	pi, err := helper.FindPlaceByPlaceName("Zonguldak")

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi == nil {
		t.Errorf("Not found\n")
	}
}

func TestFindPlaceByPlaceNameNotFound(t *testing.T) {
	helper, err := dal.NewWeatherInfoHelper()

	if err != nil {
		t.Errorf("DAL object creation problem occurred:%s\n", err.Error())
		return
	}

	if helper.ExecNonQuery("truncate table places restart identity cascade") != nil {
		t.Errorf("DB problem occurred:%s\n", err.Error())
		return
	}

	TestSavePlace(t)

	pi, err := helper.FindPlaceByPlaceName("Zongulda")

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi != nil {
		t.Errorf("Not found\n")
	}
}
