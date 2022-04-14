package models

type Employee struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Salary   float32 `json:salary`
	Func     string  `json:func`
	FilialId int     `json:filialId`
}

func GetEmployees() ([]Employee, error) {
	db, err := ConnectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `select * from "Funcionario"`

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

// func InsertClient(name, cpf string) (sql.Result, error) {
// 	db, err := ConnectDatabase()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	query := `INSERT INTO cliente (nome, "CPF") VALUES($1, $2)`

// 	res, err := db.Exec(query, name, cpf)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }

// func GetClient(cpf string) (*Client, error) {
// 	db, err := ConnectDatabase()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()

// 	query := fmt.Sprintf(`SELECT * FROM cliente WHERE "CPF"='%s'`, cpf)

// 	row, err := db.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer row.Close()

// 	row.Next()
// 	fmt.Println(row)
// 	var c Client
// 	if err := row.Scan(&c.Id, &c.Name, &c.Cpf); err != nil {
// 		fmt.Println("erro aqui")
// 		return nil, err
// 	}
// 	fmt.Println("aqui 2")
// 	return &c, nil
// }
