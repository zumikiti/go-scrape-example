package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func psqlConnect() string {
	// .env をマップ
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// port は数値で渡す必要があったので変換
	port, _ := strconv.Atoi(envs["DB_PORT"])

	// connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		envs["DB_HOST"],
		port,
		envs["DB_USER"],
		envs["DB_PASSWORD"],
		envs["DB_NAME"],
	)

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
