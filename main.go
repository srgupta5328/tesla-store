package main

import (
	"database/sql"
	"fmt"
)

//Postgres Local development credentials
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "tesla"
	dbname   = "tesla_db"
)

var storeDB *sql.DB

//InitDB initializes the database
func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")
}

func main() {
	fmt.Println("Initializing project")
}
