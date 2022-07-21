package models

import (
	"context"
	"fmt"

	"example.com/m/v2/app"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Code     string `json:"code"`
	Password string `json:"password"`
}

func (l Login) CheckUserPass(ctx context.Context) (bool, error) {
	query := `select senha from funcionario where codigo = $1 limit 1`

	row := app.GetDB().QueryRow(query, l.Code)

	var password string
	err := row.Scan(&password)
	if err != nil {
		return false, err
	}
	fmt.Println(password)
	errPassword := bcrypt.CompareHashAndPassword([]byte(password), []byte(l.Password))
	if errPassword != nil {
		return false, errPassword
	}

	return true, nil
}
