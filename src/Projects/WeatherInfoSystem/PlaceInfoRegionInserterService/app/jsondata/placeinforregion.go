package jsondata

type PlaceInfoRegion struct {
	Region string  `json:"region"`
	East   float64 `json:"east"`
	West   float64 `json:"west"`
	North  float64 `json:"north"`
	South  float64 `json:"south"`
}

func NewPlaceInfoRegion() *PlaceInfoRegion {
	return &PlaceInfoRegion{}
}
