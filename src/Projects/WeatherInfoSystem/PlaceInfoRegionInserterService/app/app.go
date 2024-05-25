package app

import (
	"PlaceInfoRegionInserterService/app/jsondata"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

const server = "http://161.97.141.113:50530"
const endPoint = server + "/api/weather/save/region"

func sendInternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Internal server error'...",
	})
}

func saveRegionServiceCallback(c *gin.Context, pi *jsondata.PlaceInfoRegion) {
	req, err := http.NewRequest("POST", endPoint, nil)
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

func saveRegionPostCallback(c *gin.Context) {
	pi := jsondata.NewPlaceInfoRegion()

	if e := c.ShouldBindJSON(pi); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	fmt.Printf("%v\n", *pi)
	//saveRegionServiceCallback(c, pi)
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func Run() {
	//Get listening port from command line argument
	engine := gin.New()
	engine.POST("/api/weather/region/save", saveRegionPostCallback) //TODO
	if e := engine.Run(); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, e.Error())
	}
}
