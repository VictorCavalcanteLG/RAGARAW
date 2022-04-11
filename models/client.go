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
