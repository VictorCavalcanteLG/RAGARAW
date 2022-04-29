package models

import (
	"database/sql"
	"fmt"

	"example.com/m/v2/app"
)

type Client struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func GetClients() ([]Client, error) {
	query := "select * from cliente"

	rows, err := app.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var c Client

		if err := rows.Scan(&c.Id, &c.Name, &c.Cpf); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return clients, nil
}

func InsertClient(name, cpf string) (sql.Result, error) {
	query := `INSERT INTO cliente (nome, "CPF") VALUES($1, $2)`

	res, err := app.GetDB().Exec(query, name, cpf)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetClient(cpf string) (*Client, error) {
	query := fmt.Sprintf(`SELECT * FROM cliente WHERE "CPF"='%s'`, cpf)

	row, err := app.GetDB().Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()

	var c Client
	if err := row.Scan(&c.Id, &c.Name, &c.Cpf); err != nil {
		return nil, err
	}

	return &c, nil
}
