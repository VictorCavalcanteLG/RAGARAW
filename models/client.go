package models

import "database/sql"

func GetClients() (*sql.Rows, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}

	query := "select * from cliente"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func InsertClient(name, cpf string) (sql.Result, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO cliente (nome, "CPF") VALUES($1, $2)`

	res, err := db.Exec(query, name, cpf)
	if err != nil {
		return nil, err
	}

	return res, nil
}
