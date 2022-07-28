package models

import (
	"database/sql"

	"example.com/m/v2/app"
)

type Store struct {
	IdStore int
	Name    string `json:"store_name"`
	City    string `json:"city"`
}

func InsertStore(s Store) (sql.Result, error) {
	query := `insert into filial (nome, cidade) values ($1, $2)`

	res, err := app.GetDB().Exec(query, s.Name, s.City)
	if err != nil {
		return nil, err
	}

	return res, nil
}
