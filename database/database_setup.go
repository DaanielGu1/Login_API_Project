package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func errorDatabaseAccess() {
	var value string
	fmt.Println("Create a new database file? [Y/n]")
	fmt.Scan(&value)
	if (value=="Y") {
		createFileForError()
	}
}

func createFileForError() {
	var filePath string
	var err error

	fmt.Println("Type the path to where your file is going to be created: ")
	fmt.Scan(&filePath)

	filePath = filePath + "/database.db"

	err = os.WriteFile(filePath, nil, 0644)  
	if err != nil {
		log.Fatalf("Error trying to find path %v", err)
	}
}


func CreateDatabase() {

	var err error
	

	db, err = sql.Open("sqlite3", "database.db?_foreign_keys=on")
	if err != nil {
		fmt.Printf("Error trying to access the database: %v", err)
		errorDatabaseAccess()
		db, err = sql.Open("sqlite3", "database.db?_foreign_keys=on")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error trying to interact with the database: %v", err)
	}
	//"DEFAULT CURRENT_TIMESTAMP" may cause divergence depending on the user's computer clock
	table_query, err := os.ReadFile("./SQL/create_tables.sql")
	if err != nil {
		log.Fatalf("Error trying to read the file create_tables.sql in folder SQL: %v", err)
	}

	_, err = db.Exec(string(table_query))
	if err != nil {
		log.Fatalf("Error trying to build the database tables: %v", err)
	}
}