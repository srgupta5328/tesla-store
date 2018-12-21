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

func main() {
	fmt.Println("Initializing project")
}
