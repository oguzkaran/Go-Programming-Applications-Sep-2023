package app

import (
	"SampleTimeServiceApp/app/data/dal"
	"SampleTimeServiceApp/app/data/repository/entity"
	"SampleTimeServiceApp/app/jsondata"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var helper *dal.TimeServiceHelper

func initDB() {
	h, err := dal.NewTimeServiceHelper()
	if err != nil {
		fmt.Printf("Can not connecto to db:%s\n", err.Error())
		os.Exit(1)
	}

	helper = h
}

func saveTimeClientInfo(ci *jsondata.ClientInfo) {
	var cie = &entity.TimeClientInfo{Host: ci.Host, Name: ci.Name, DateTime: ci.DateTime}
	helper.SaveTimeClientInfo(cie)
}

func saveDateClientInfo(ci *jsondata.ClientInfo) {
	var die = &entity.DateClientInfo{Host: ci.Host, Name: ci.Name, Date: ci.DateTime}
	helper.SaveDateClientInfo(die)
}

func timeHandler(c *gin.Context) {
	value := c.Request.FormValue("name")

	if value != "" {
		now := time.Now()
		ci := jsondata.NewClientInfo(c.Request.RemoteAddr, "Hello "+value, now.Format("02/01/2006 15:04:05"))
		saveTimeClientInfo(ci)
		c.IndentedJSON(http.StatusOK, ci)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameter 'name' required!..."})
	}

}

func dateHandler(c *gin.Context) {
	value := c.Request.FormValue("name")

	if value != "" {
		now := time.Now()
		ci := jsondata.NewClientInfo(c.Request.RemoteAddr, "Hello "+value, now.Format("02/01/2006"))
		saveDateClientInfo(ci)
		c.IndentedJSON(http.StatusOK, ci)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameter 'name' required!..."})
	}
}

func birthDateHandler(c *gin.Context) {
	value := c.Param("name")
	if value != "" {
		c.JSON(http.StatusOK, gin.H{"message": "Will be available later!..."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Parameter 'name' required!..."})
	}
}

func defaultHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "Not Found!..."})
}

func checkArguments() {
	if len(os.Args) != 2 {
		fmt.Printf("Wrong number of arguments!...")
		os.Exit(1)
	}
}

func addEndPoints(engine *gin.Engine) {
	engine.GET("/time", timeHandler)
	engine.GET("/date", dateHandler)
	engine.GET("/date/birth", birthDateHandler)
	engine.GET("/", defaultHandler)
}

func startServer(engine *gin.Engine) {
	err := engine.Run(":" + os.Args[1])
	if err != nil {
		fmt.Printf("Message:%s\n", err.Error())
	}
}

func Run() {
	checkArguments()
	initDB()
	engine := gin.New()
	addEndPoints(engine)
	startServer(engine)
}
