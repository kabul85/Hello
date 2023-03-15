package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
			"method":  "This is post method",
		})
	})

	r.Run()
}
