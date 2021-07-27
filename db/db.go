package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConexaoDB() *sql.DB {
	// conexao := "user=postgres dbname=alura_loja password=meli1705 host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", conexao)
	// if err != nil {
	// 	panic(err.Error())
	// }

	db, err := sql.Open("mysql", "root:@/alura_loja")
	if err != nil {
		panic(err.Error())
	}

	return db
}
