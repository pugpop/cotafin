package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Contact struct {
	ContactId  int
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
	ContactId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	FirstName1 TEXT
	LastName1 TEXT
	FirstName2 TEXT
	LastName2 TEXT
	Address TEXT
	City TEXT
	State TEXT
	Zipcode TEXT
	Phone TEXT
	Notes TEXT
	);`

	_, err = DB.Exec(sqlStatement)
	check(err)

	sqlStatement = `
	CREATE TABLE IF NOT EXISTS xaction (
	XactionID INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	SpecialCollection INTEGER
	Xamount REAL
	);`

	_, err = DB.Exec(sqlStatement)
	check(err)

}

func main() {
	initDB()
	//defer DB.Close()
	fmt.Println("done")
}
