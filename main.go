package main

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.ListClients())

	r.Run()
}
