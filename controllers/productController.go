package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func InsertProduct() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var product models.Product

		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		res, err := models.InsertProduct(product)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		rowsAffected, _ := res.RowsAffected()
		ctx.JSON(http.StatusCreated, gin.H{"rows affected": rowsAffected})
	}
}
