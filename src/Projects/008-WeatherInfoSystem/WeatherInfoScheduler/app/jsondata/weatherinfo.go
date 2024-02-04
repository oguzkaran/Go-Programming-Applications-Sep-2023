package jsondata

type WeatherInfo struct {
	WeatherObservations []struct {
		Lng              float64 `json:"lng"`
		Observation      string  `json:"observation"`
		Icao             string  `json:"ICAO"`
		Clouds           string  `json:"clouds"`
		DewPoint         string  `json:"dewPoint"`
		Datetime         string  `json:"string"`
		Temperature      string  `json:"temperature"`
		Humidity         float64 `json:"humidity"`
		StationName      string  `json:"stationName"`
		WeatherCondition string  `json:"weatherCondition"`
		WindDirection    float64 `json:"windDirection"`
		WindSpeed        string  `json:"windSpeed"`
		Lat              float64 `json:"lat"`
		CloudsCode       string  `json:"cloudsCode,omitempty"`
	} `json:"weatherObservations"`
}
