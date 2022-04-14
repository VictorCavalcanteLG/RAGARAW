package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func ListClients() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		modelClients, err := models.GetClients()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		ctx.JSON(http.StatusOK, modelClients)
	}
}

func InsertClient() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var c models.Client
		ctx.ShouldBindJSON(&c)

		res, err := models.InsertClient(c.Name, c.Cpf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		rowsAffected, _ := res.RowsAffected()
		ctx.JSON(http.StatusCreated, gin.H{"rows affected": rowsAffected})
	}
}

func GetClient() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var client models.Client
		ctx.ShouldBindJSON(&client)
		fmt.Println(client.Cpf)
		modelClient, err := models.GetClient(client.Cpf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Println(modelClient)

		ctx.JSON(http.StatusOK, *modelClient)
	}
}
