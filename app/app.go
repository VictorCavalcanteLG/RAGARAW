package app

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 1234
	user     = "postgres"
	password = "postgres"
	dbname   = "ragaraw"
)

var db *sql.DB

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	d, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}

	err = d.Ping()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic(err)
	}

	db = d
}

func GetDB() *sql.DB {
	return db
}
