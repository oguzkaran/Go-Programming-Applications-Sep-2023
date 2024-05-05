package app

import (
	"PlaceInfoInserter/app/converter"
	"PlaceInfoInserter/app/data/entity"
	"PlaceInfoInserter/app/jsondata"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func savePlaceCallback(c *gin.Context, db *gorm.DB) {
	var pi jsondata.PlaceInfoSaveDTO

	if e := c.ShouldBindJSON(&pi); e == nil {
		db.Create(converter.ToPlaceInfoSave(&pi))
		fmt.Println(pi)
		c.JSON(http.StatusCreated, pi)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": e.Error(),
		})
	}
}

func addEndPoints(e *gin.Engine, db *gorm.DB) {
	e.POST("/api/places/save", func(context *gin.Context) {
		savePlaceCallback(context, db)
	})
}

func initDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=csd1993 sslmode=disable dbname=gs23_weatherinfosystemdb_dev")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "database initialization problem:%s\n", err.Error())
		os.Exit(1)
	}

	db.AutoMigrate(&entity.PlaceInfo{})

	return db
}

func Run() {
	db := initDB()
	e := gin.New()
	addEndPoints(e, db)
	if err := e.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
