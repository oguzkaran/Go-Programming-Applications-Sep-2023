package converter

import (
	"PlaceInfoInserter/app/data/entity"
	"PlaceInfoInserter/app/jsondata"
)

func ToPlaceInfoSave(pi *jsondata.PlaceInfoSaveDTO) *entity.PlaceInfo {
	return &entity.PlaceInfo{Name: pi.Name, Latitude: pi.Latitude, Longitude: pi.Longitude}
}
