package models

import (
	"database/sql"
	"fmt"
)

type Employee struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Salary   float32 `json:"salary"`
	Func     string  `json:"func"`
	FilialId int     `json:"filialId"`
}

func GetEmployees() ([]Employee, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `select * from funcionario`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var c Employee

		if err := rows.Scan(&c.Id, &c.Name, &c.Code, &c.Salary, &c.Func, &c.FilialId); err != nil {
			return nil, err
		}

		employees = append(employees, c)
	}

	return employees, nil
}

func InsertEmployee(emp Employee) (sql.Result, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `INSERT INTO funcionario (nome, codigo, salario, funcao, filial_id)
	VALUES($1, $2, $3, $4, $5)`

	res, err := db.Exec(query, emp.Name, emp.Code, emp.Salary, emp.Func, emp.FilialId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetEmployee(code string) (*Employee, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := fmt.Sprintf(`select * from funcionario where codigo='%s'`, code)

	row, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	fmt.Println(row)
	var e Employee
	if err := row.Scan(&e.Id, &e.Name, &e.Code, &e.Salary, &e.Func, &e.FilialId); err != nil {
		return nil, err
	}

	return &e, nil
}
