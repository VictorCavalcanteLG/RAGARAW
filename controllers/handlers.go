package controllers

import (
	"fmt"

	"example.com/m/v2/globals"
	"example.com/m/v2/helpers"
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Please logout first"})
			return
		}

		var l Login
		err := c.ShouldBindJSON(&l)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		username := l.Password
		password := l.Username
		fmt.Println(username, password)

		if helpers.EmptyUserPass(username, password) {
			c.JSON(http.StatusBadRequest, gin.H{"content": "Parameters can't be empty"})
			return
		}

		if !helpers.CheckUserPass(username, password) {
			c.JSON(http.StatusUnauthorized, gin.H{"content": "Incorrect username or password"})
			return
		}

		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"content": "Failed to save session"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
}

func LogoutGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
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

		c.Redirect(http.StatusMovedPermanently, "/")
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
