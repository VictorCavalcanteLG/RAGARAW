package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func ListEmployees() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		modelEmployees, err := models.GetEmployees()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		ctx.JSON(http.StatusOK, modelEmployees)
	}
}

func InsertEmployee() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var e models.Employee
		ctx.ShouldBindJSON(&e)

		res, err := models.InsertEmployee(e)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		rowsAffected, _ := res.RowsAffected()
		ctx.JSON(http.StatusCreated, gin.H{"rows affected": rowsAffected})
	}
}

func GetEmployee() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var employee models.Employee
		ctx.ShouldBindJSON(&employee)

		modelEmployee, err := models.GetEmployee(employee.Code)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		ctx.JSON(http.StatusOK, modelEmployee)
	}
}
