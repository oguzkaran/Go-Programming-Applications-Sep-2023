package converter

import (
	"PlaceInfoInserter/app/data/entity"
	"PlaceInfoInserter/app/jsondata"
)

func ToPlaceInfoSave(pi *jsondata.PlaceInfoSaveDTO) *entity.PlaceInfoLocation {
	return &entity.PlaceInfoLocation{Name: pi.Name, Latitude: pi.Latitude, Longitude: pi.Longitude}
}
