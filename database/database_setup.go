package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func CreateDatabase() {

	var err error
	

	db, err = sql.Open("sqlite3", "database.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Error trying to access the database: %v", err)
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
