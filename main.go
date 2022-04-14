package main

import (
	"example.com/m/v2/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/clients", controllers.ListClients())
	r.POST("/client", controllers.InsertClient())
	r.GET("/client", controllers.GetClient())

	r.Run(":3000")
}
