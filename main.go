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
	r.GET("/employees", controllers.ListEmployees())
	r.POST("/employee", controllers.InsertEmployee())
	r.GET("/employee", controllers.GetEmployee())

	r.Run(":3000")
}
