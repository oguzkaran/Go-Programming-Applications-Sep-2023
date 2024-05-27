package app

import (
	"PlaceInfoRegionInserterService/app/jsondata"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const server = "http://161.97.141.113:50530"
const endPoint = server + "/api/weather/places/region/save"

func sendInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

func saveRegionServiceCallback(c *gin.Context, pi *jsondata.PlaceInfoRegion) {
	data, err := json.Marshal(*pi)
	if err != nil {
		sendInternalServerError(c, "Internal server error")
		return
	}

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(data))
	if err != nil {
		sendInternalServerError(c, "Internal server error")
		return
	}

	req.Header.Add("Content-Type", "application/json")

	client := http.Client{Timeout: 20 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		sendInternalServerError(c, "Internal server error")
		return
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode == http.StatusBadRequest {
		sendInternalServerError(c, "Bad request!...")
		return
	}

	if res.StatusCode != http.StatusOK {
		sendInternalServerError(c, "Internal server error")
		return
	}

	data, err = io.ReadAll(res.Body)

	if err != nil {
		sendInternalServerError(c, "Internal server error")
		return
	}

	err = json.Unmarshal(data, pi)

	if err != nil {
		sendInternalServerError(c, "Internal server error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func saveRegionPostCallback(c *gin.Context) {
	pi := jsondata.NewPlaceInfoRegion()

	if e := c.ShouldBindJSON(pi); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": e.Error()})
		return
	}

	fmt.Printf("%v\n", *pi)
	saveRegionServiceCallback(c, pi)
}

func Run() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "wrong number of arguments")
		os.Exit(1)
	}
	port, e := strconv.Atoi(os.Args[1])

	if e != nil {
		_, _ = fmt.Fprintf(os.Stderr, "invalid port value")
		os.Exit(1)
	}

	engine := gin.New()
	engine.POST("/api/weather/region/save", saveRegionPostCallback)
	if e := engine.Run(fmt.Sprintf(":%d", port)); e != nil {
		_, _ = fmt.Fprintf(os.Stderr, e.Error())
	}
}
