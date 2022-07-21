package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func CreateEmployee() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var e models.Employee

		err := ctx.ShouldBindJSON(&e)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}

		if e.Password == "" {
			fmt.Printf("password not defined")
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": "password not defined"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		e.Password = string(hashedPassword)
		res, err := models.InsertEmployee(e)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
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
