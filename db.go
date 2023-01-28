package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "sail"
	password = "password"
	dbname   = "kabu"
)

func psqlConnect() string {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return psqlconn
}

func saveData(data Result) {
	psqlconn := psqlConnect()

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// insert
	insertStmt := `insert into "prices"("price", "per", "pbr") values($1, $2, $3)`
	_, e := db.Exec(insertStmt, data.price, data.per, data.pbr)
	CheckError(e)

	fmt.Println("Inserted!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
