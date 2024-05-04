/*
Sınıf Çalışması: Aşağıdaki GET servise erişerek bir posta koduna ilişkin JSON veriyi elde ediniz. Uygulamada yalnızca
posta kodu alınacaktır. Diğer parametreler URL'de görüldüğü şekilde olacaktır
http://api.geonames.org/postalCodeSearchJSON?formatted=true&postalcode=67000&maxRows=10&username=csystem&country=tr
*/

/*
Sınıf Çalışması: Uygulamayı posta kodu bilgisini yayınlayan ve aşağıdaki gibi çalışam bir REST servis durumuna getiriniz.
	Açıklamalar:
		- Sorgulanan posta kodu bilgisi için önce ddl.sql içerisinde yaratılan veritabanına bakılacak ve eğer veritabanında
		posta kodu bulunuyorsa bilgiler veritabanından alınacaktır. Bulunmuyorsa geonames'e gidilerek ilgili posta kodunda
		ilişkin bilgiler elde edilecek ve veritabanına eklenecektir.

		- Veritabanında posta kodlarının sorgulama tarih zamanları da tutulmaktadır.

		- Sorgulama için maxRows değeri consumer'dan istenmeyecek default olarak 10 değeriyle alınacaktır.

		- Sorgulama sayısı da istenen ayrı bir Handler yazılacak ve buradan istekte bulunulduğunda hep geonames'e
		erişilecektir

		- Projeyi Repository, Service ve Application katmanı olacak şekilde tasarlayınız. Gerekirse DAL katmanı da
		ekleyebilirsiniz

		- Veritabanı erişimi için önce lib/pq kullanınız. Daha sonra gorm ile yeniden implemente edebilirsiniz

		- Servislerin URL bilgilerini dilediğiniz gibi belirleyebilirsiniz
*/

package app

import (
	"PostalCodeSearchService/app/jsondata"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const server = "http://api.geonames.org"
const postalCodeSearcEndPoint = "/postalCodeSearchJSON"
const postalCodeSearchURL = server + postalCodeSearcEndPoint

func sendInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error'...",
	})
}
func postalCodeSearchCallback(c *gin.Context, postalCode int, server string) {
	url := fmt.Sprintf("%s?formatted=true&postalcode=%d&maxRows=10&username=csystem&country=tr", postalCodeSearchURL, postalCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		sendInternalServerError(c)
		return
	}

	client := http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		sendInternalServerError(c)
		return
	}
	pi := jsondata.PostalCodeInfo{}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		sendInternalServerError(c)
		return
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		sendInternalServerError(c)
		return
	}

	err = json.Unmarshal(data, &pi)

	if err != nil {
		sendInternalServerError(c)
		return
	}
	c.IndentedJSON(http.StatusOK, pi)
}

func postalCodeGetCallback(c *gin.Context) {
	codeStr := c.Request.FormValue("code")
	if codeStr != "" {
		code, e := strconv.Atoi(codeStr)
		if e == nil {
			postalCodeSearchCallback(c, code, "")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Postal code must be a numeric value!...",
			})
		}

	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Parameter 'code' required!...",
		})
	}
}

func Run() {
	engine := gin.New()
	engine.GET("/api/postalcode", postalCodeGetCallback)
	if e := engine.Run(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, e.Error())
	}
}
