package controllers

import (
	"database/sql"
	"fmt"

	"example.com/m/v2/models"
)

func GetClients() (*sql.Rows, error) {
	db, err := models.ConnectDatabase()
	if err != nil {
		return nil, err
	}

	query := "select * from cliente"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	return rows, nil
}

func ListClients() {
	rows, err := GetClients()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	for rows.Next() {
		var (
			id        int
			nome, cpf string
		)

		if err := rows.Scan(&id, &nome, &cpf); err != nil {
			fmt.Printf("err: %v\n", err)
		}

		fmt.Println(id, nome, cpf)
	}
}
