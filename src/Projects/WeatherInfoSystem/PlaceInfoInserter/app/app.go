package app

import (
	"PlaceInfoInserter/app/jsondata"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func savePlaceCallback(c *gin.Context) {
	var pi jsondata.PlaceInfoSave

	if e := c.ShouldBindJSON(&pi); e == nil {
		//save info to database
		fmt.Println(pi)
		c.JSON(http.StatusCreated, pi)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
	}
}

func addEndPoints(e *gin.Engine) {
	e.POST("/api/places/save", savePlaceCallback)
}

func Run() {
	e := gin.New()
	addEndPoints(e)
	if err := e.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
