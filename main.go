package main

import (
	"example.com/m/v2/app"
	"example.com/m/v2/controllers"
	"example.com/m/v2/globals"
	"example.com/m/v2/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	app.ConnectDatabase()

	r := gin.Default()

	r.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := r.Group("/")
	// routes.PublicRoutes(public)
	public.GET("/login", controllers.LoginGetHandler())
	public.POST("/login", controllers.LoginPostHandler())
	public.GET("/", controllers.IndexGetHandler())

	private := r.Group("/")
	private.Use(middleware.AuthRequired)
	// routes.PublicRoutes(private)

	// r.GET("/clients", controllers.ListClients())
	// r.POST("/client", controllers.InsertClient())
	// r.GET("/client", controllers.GetClient())
	private.GET("/employees", controllers.ListEmployees())
	private.GET("/logout", controllers.LogoutGetHandler())
	private.POST("/employee", controllers.CreateEmployee())
	// r.GET("/employee", controllers.GetEmployee())

	r.Run(":3000")
}
