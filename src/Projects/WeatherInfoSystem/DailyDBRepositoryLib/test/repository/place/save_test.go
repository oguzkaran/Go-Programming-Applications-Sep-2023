package place

import (
	"WeatherInfoScheduler/data/entity"
	repositoryTest "WeatherInfoScheduler/data/repository"
	"WeatherInfoScheduler/test/repository"
	"testing"
)

func TestSavePlace(t *testing.T) {
	db, err := repository.InitDB()

	if err != nil {
		t.Errorf("DB problem occurred:%s\n", err.Error())
		return
	}

	pr := repositoryTest.NewPlaceRepository(db)

	if pr == nil {
		t.Errorf("DB problem occurred\n")
		return
	}
	placeInfo := &entity.Place{Code: 67, PlaceName: "Zonguldak", Country: "TR", North: 23.4, South: 34.5, East: 56.7, West: 45.7}

	err = pr.Save(placeInfo)

	if err != nil {
		t.Errorf("Save problem occurred:%s\n", err.Error())
		return
	}
}
