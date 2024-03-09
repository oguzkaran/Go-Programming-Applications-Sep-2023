package app

import (
	"WeatherInfoScheduler/app/jsondata"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const server = "http://api.geonames.org"
const weatherInfoEndPoint = "/weatherJSON"
const weatherInfoURL = server + weatherInfoEndPoint

func weatherInfoCallback(north, south, east, west float64) (int, string) {
	url := fmt.Sprintf("%s?formatted=true&north=%f&south=%f&east=%f&west=%f&username=csystem",
		weatherInfoURL, north, south, east, west)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error in request:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	client := http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	wi := jsondata.WeatherInfoView{}

	defer res.Body.Close()

	if err != nil {
		fmt.Printf("Client error:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	if res.StatusCode != http.StatusOK {
		return res.StatusCode, ""
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Data error:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	err = json.Unmarshal(data, &wi)

	if err != nil {
		fmt.Printf("Data Unmarshal error:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	for _, wo := range wi.WeatherObservations {
		fmt.Printf("%s", wo.Observation)
	}

	return res.StatusCode, string(data)
}

func readPlaceInfo(prompt string) (float64, float64, float64, float64) {
	var north, south, east, west float64

	fmt.Print(prompt)
	fmt.Scanf("%f%f%f%f", &north, south, east, west)

	return north, south, east, west
}

func Run() {
	for {
		status, message := weatherInfoCallback(readPlaceInfo("Input place information:"))
		if status == http.StatusOK {
			fmt.Println(message)
		} else {
			fmt.Printf("Status Code:%d\n", status)
		}
	}
}
