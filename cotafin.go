package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type CotaFIN struct {
	ID         int
	FirstName1 string
	LastName1  string
	FirstName2 string
	LastName2  string
	Address    string
	City       string
	State      string
	Zipcode    string
	Phone      string
	Notes      string
}

var DB *sql.DB

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() {
	DB, err := sql.Open("sqlite3", "./cotafin.db")
	check(err)

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS contacts (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	firstName1 TEXT
	lastName1 TEXT
	firstName2 TEXT
	lastName2 TEXT
	address TEXT
	city TEXT
	state TEXT
	zipcode TEXT
	phone TEXT
	notes TEXT
	);`

	_, err = DB.Exec(sqlStatement)
	check(err)
}

func main() {
	initDB()
	//defer DB.Close()
	fmt.Println("done")
}
