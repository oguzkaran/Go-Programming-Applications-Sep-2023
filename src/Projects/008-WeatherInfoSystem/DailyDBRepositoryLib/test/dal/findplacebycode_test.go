package dal

import (
	"WeatherInfoScheduler/data/dal"
	"testing"
)

func TestFindPlaceByCode(t *testing.T) {
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

	pi, err := helper.FindPlaceByCode(67)

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi == nil {
		t.Errorf("Not found\n")
	}
}

func TestFindPlaceByCodeNotFound(t *testing.T) {
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

	pi, err := helper.FindPlaceByCode(37)

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi != nil {
		t.Errorf("Found\n")
	}
}
