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
	c_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	c_first_name1 TEXT,
	c_last_name1 TEXT,
	c_first_name2 TEXT,
	c_last_name2 TEXT,
	c_address TEXT,
	c_city TEXT,
	c_state TEXT,
	c_zipcode TEXT,
	c_phone TEXT,
	c_notes TEXT,
	c_xaction_id INTEGER,
	FOREIGN KEY (c_xaction_id) REFERENCES xaction(x_id)
	);`

	_, err = DB.Exec(sqlStatement)
	check(err)

	sqlStatement = `
	CREATE TABLE IF NOT EXISTS xaction (
	x_iD INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	x_special_collection INTEGER,
	x_miscellaneous INTEGER,
	x_amount REAL
	);`

	_, err = DB.Exec(sqlStatement)
	check(err)

}

func main() {
	initDB()
	//defer DB.Close()
	fmt.Println("done")
}
