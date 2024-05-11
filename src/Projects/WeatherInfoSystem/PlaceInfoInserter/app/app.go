package app

import (
	"PlaceInfoInserter/app/converter"
	"PlaceInfoInserter/app/data/repository"
	"PlaceInfoInserter/app/jsondata"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
)

func savePlaceCallback(c *gin.Context, r *repository.PlaceInfoRepository) {
	var pi jsondata.PlaceInfoSaveDTO

	if e := c.ShouldBindJSON(&pi); e == nil {
		r.Save(converter.ToPlaceInfoSave(&pi))
		fmt.Println(pi)
		c.JSON(http.StatusCreated, pi)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
	}
}

func addEndPoints(e *gin.Engine, r *repository.PlaceInfoRepository) {
	e.POST("/api/places/save", func(context *gin.Context) {
		savePlaceCallback(context, r)
	})
}

func checkArguments(length int) {
	if length != len(os.Args) {
		_, _ = fmt.Fprintf(os.Stderr, "wrong number of arguments\n")
		os.Exit(1)
	}
}

func Run() {
	checkArguments(2)
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Invalid server port!...")
		os.Exit(1)
	}

	addr := fmt.Sprintf(":%d", port)
	e := gin.New()

	r := repository.NewPlaceInfoRepository()
	if r == nil {
		_, _ = fmt.Fprintf(os.Stderr, "Repository prblem occurred!...")
	}
	addEndPoints(e, r)
	if err := e.Run(addr); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
