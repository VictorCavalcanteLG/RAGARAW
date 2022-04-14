package models

import (
	"database/sql"
	"fmt"
)

type Client struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

func GetClients() ([]Client, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "select * from cliente"

	rows, err := db.Query(query)
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
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `INSERT INTO cliente (nome, "CPF") VALUES($1, $2)`

	res, err := db.Exec(query, name, cpf)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetClient(cpf string) (*Client, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf(`SELECT * FROM cliente WHERE "CPF"='%s'`, cpf)

	row, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	fmt.Println(row)
	var c Client
	if err := row.Scan(&c.Id, &c.Name, &c.Cpf); err != nil {
		fmt.Println("erro aqui")
		return nil, err
	}
	fmt.Println("aqui 2")
	return &c, nil
}
