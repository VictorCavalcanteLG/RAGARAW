package main

import (
	"net/http"

	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	controllers.ListClients()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
