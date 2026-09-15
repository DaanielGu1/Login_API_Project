package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func CreateDatabase() *sql.DB {

	var err error
	

	db, err = sql.Open("sqlite3", "database.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("Error trying to access the database: %v", err)
	}

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
	return db
}

func CreateNewUser(db *sql.DB, first_name, last_name, birthday, password string) (int64, error) {
	new_user := `INSERT INTO users (first_name, last_name, birthday, password)
	VALUES (?, ?, ?, ?)`

	result, err := db.Exec(new_user, first_name, last_name, birthday, password)
	if err != nil {
		log.Fatalf("Error trying to insert data to a new user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error trying to get user's ID: %v", err)
	}

	return id, err
}

func CreateNewGroup(db *sql.DB, group_name string) (int64, error) {
	new_group := `INSERT INTO groups (group_name)
	VALUES (?)`

	result, err := db.Exec(new_group, group_name)
	if err != nil {
		log.Fatalf("Error trying to insert data to a new user: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("Error trying to get user's ID: %v", err)
	}

	return id, err
}