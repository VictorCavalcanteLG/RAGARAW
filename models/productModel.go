package models

import (
	"database/sql"

	"example.com/m/v2/app"
)

type Product struct {
	Product string `json:"product_name"`
	Code    string `json:"code"`
}

func InsertProduct(p Product) (sql.Result, error) {
	query := `insert into produto (nome, codigo) values ($1, $2)`

	res, err := app.GetDB().Exec(query, p.Product, p.Code)
	if err != nil {
		return nil, err
	}

	return res, nil
}
