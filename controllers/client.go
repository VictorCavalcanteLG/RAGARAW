package controllers

import (
	"fmt"
	"net/http"

	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

type Client struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func ListClients() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		rows, err := models.GetClients()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		defer rows.Close()

		var clients []Client
		for rows.Next() {
			var c Client

			if err := rows.Scan(&c.Id, &c.Name, &c.Cpf); err != nil {
				fmt.Printf("err: %v\n", err)
			}

			clients = append(clients, c)
		}

		ctx.JSON(http.StatusOK, clients)
	}
}

func InsertClient() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fmt.Println("teste")
		res, err := models.InsertClient("kk", "00000000000")
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("res: %v\n", res)
	}
}
