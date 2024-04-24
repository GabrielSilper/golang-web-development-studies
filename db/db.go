package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDb() *sql.DB {
	uri := "postgresql://root:password@localhost/loja?sslmode=disable"
	db, err := sql.Open("postgres", uri)

	if err != nil {
		panic(err.Error())
	}

	return db
}
