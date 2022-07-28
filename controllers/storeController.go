package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func InsertStore() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var store models.Store

		err := ctx.ShouldBindJSON(&store)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		res, err := models.InsertStore(store)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		rowsAffected, _ := res.RowsAffected()
		ctx.JSON(http.StatusCreated, gin.H{"rows affected": rowsAffected})
	}
}
