/*
------------------------------------------------------------------------------------------------------------------------



------------------------------------------------------------------------------------------------------------------------
*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func greetingCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
		"company": "CSD",
	})
}

func defaultCallback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Default",
		"company": "CSD",
	})
}

func main() {
	e := gin.New()
	e.GET("/greeting", greetingCallback)
	e.GET("/", defaultCallback)
	if err := e.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}
}
