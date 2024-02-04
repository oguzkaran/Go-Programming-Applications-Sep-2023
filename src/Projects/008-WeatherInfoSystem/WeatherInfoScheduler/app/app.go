/*
Sınıf Çalışması: Elimizde bulunan yer bilgisine göre periyodik olarak Geonames'den hava durumu bilgilerini çeken
bir scheduler job yazınız. Bilgilerin sayısının önemi yoktur
*/
package app

import (
	"PostalCodeSearchService/app/jsondata"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const server = "http://api.geonames.org"
const weatherInfoEndPoint = "/weatherJSON"
const weatherInfoURL = server + weatherInfoEndPoint

func postalCodeSearchCallback(postalCode int, server string) (int, string) {
	url := fmt.Sprintf("%s?formatted=true&north=%f&south=%f&east=%f&west=%f&username=csystem", weatherInfoURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error in request:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	client := http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req) //Response'un kapatılması gerekir. İleride defer function'lar yapacağız
	pi := jsondata.PostalCodeInfo{}

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

	err = json.Unmarshal(data, &pi)

	if err != nil {
		fmt.Printf("Data Unmarshalerror:%s", err.Error())
		return http.StatusInternalServerError, ""
	}

	for _, pc := range pi.PostalCodes {
		fmt.Println(pc.PlaceName)
	}

	return res.StatusCode, string(data)
}

func readPostalCode(prompt string) int {
	var code int

	fmt.Print(prompt)
	fmt.Scanf("%d", &code)

	return code
}

func Run() {
	for {
		code := readPostalCode("Input postal code:")
		if code <= 0 {
			break
		}
		status, message := postalCodeSearchCallback(code, server)

		if status == http.StatusOK {
			fmt.Println(message)
		} else {
			fmt.Printf("Status Code:%d\n", status)
		}
	}
}
