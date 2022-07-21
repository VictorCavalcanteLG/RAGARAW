package controllers

import (
	"fmt"

	"example.com/m/v2/globals"
	"example.com/m/v2/helpers"
	"example.com/m/v2/models"
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func LoginPostHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get(globals.Userkey)
		if user != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}

		var userLogin models.Login
		err := ctx.ShouldBindJSON(&userLogin)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if helpers.EmptyUserPass(userLogin.Code, userLogin.Password) {
			ctx.JSON(http.StatusBadRequest, gin.H{"content": "Parameters can't be empty"})
			return
		}

		userExists, err := userLogin.CheckUserPass(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"content": "Internal error"})
			return
		}

		if !userExists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"content": "Incorrect code or password"})
			return
		}

		session.Set(globals.Userkey, userLogin.Code)
		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}

		ctx.JSON(http.StatusMovedPermanently, gin.H{"content": "LOGGED"})
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		user := session.Get(globals.Userkey)
		log.Println("logging out user:", user)

		if user == nil {
			log.Println("Invalid session token")
			return
		}

		session.Delete(globals.Userkey)

		if err := session.Save(); err != nil {
			log.Println("Failed to save session:", err)
			return
		}

		ctx.JSON(http.StatusMovedPermanently, gin.H{"content": "LOGOUT"})
	}
}

func IndexGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H{
			"content": "This is an index page...",
			"user":    user,
		})
	}
}

func DashboardGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		c.JSON(http.StatusOK, gin.H{
			"content": "This is a dashboard",
			"user":    user,
		})
	}
}
