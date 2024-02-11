package place

import (
	repositoryTest "WeatherInfoScheduler/data/repository"
	"WeatherInfoScheduler/test/repository"
	"testing"
)

func TestFindPlaceByPlaceName(t *testing.T) {
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

	result := db.Exec("truncate table places restart identity cascade")
	if result.Error != nil {
		t.Errorf("DB problem occurred:%s\n", result.Error.Error())
		return
	}

	TestSavePlace(t)

	pi, err := pr.FindByPlaceName("Zonguldak")

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi == nil {
		t.Errorf("Not found\n")
	}
}

func TestFindPlaceByPlaceNameNotFound(t *testing.T) {
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

	result := db.Exec("truncate table places restart identity cascade")
	if result.Error != nil {
		t.Errorf("DB problem occurred:%s\n", result.Error.Error())
		return
	}

	TestSavePlace(t)

	pi, err := pr.FindByPlaceName("Zonguldk")

	if err != nil {
		t.Errorf("DB Problem occured:%s\n", err.Error())
		return
	}

	if pi != nil {
		t.Errorf("Found\n")
	}
}
