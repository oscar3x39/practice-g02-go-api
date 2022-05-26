package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main () {

	// new instance
	r := gin.New()
	// register middleware
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "頁面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "404 not found",
			})
		}
	})

	r.Run(":8000")
}