package jsondata

type PostalCodeInfo struct {
	PostalCodes []struct {
		AdminCode1  string  `json:"adminCode1"`
		AdminName2  string  `json:"adminName2"`
		Longitude   float64 `json:"lng"`
		CountryCode string  `json:"countryCode"`
		PostalCode  string  `json:"postalCode"`
		AdminName1  string  `json:"adminName1"`
		ISO31662    string  `json:"ISO3166-2"`
		PlaceName   string  `json:"placeName"`
		Latitude    float64 `json:"lat"`
	} `json:"postalCodes"`
}
