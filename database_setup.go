package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {

	var ERROR error

	db, ERROR = sql.Open("sqlite3", "database.db?_foreign_keys=on")
	if ERROR != nil {
		log.Fatalf("Error trying to access the database: %v", ERROR)
	}
	defer db.Close()

	ERROR = db.Ping()
	if ERROR != nil {
		log.Fatalf("Error trying to interact with the database: %v", ERROR)
	}
	//"DEFAULT CURRENT_TIMESTAMP" may cause divergence depending on the user's computer clock
	table_query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username CHAR(50) UNIQUE,
		first_name CHAR(50) NOT NULL,
		last_name CHAR(50) NOT NULL,
		birthday DATE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
		password VARCHAR(255) NOT NULL,
		email CHAR(50),
		account_status BOOLEAN DEFAULT 1
	);

	CREATE TABLE IF NOT EXISTS groups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		group_name CHAR(50) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		description VARCHAR(255)
	);

	CREATE TABLE IF NOT EXISTS group_members (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		group_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
	);`
	_, ERROR = db.Exec(table_query)
	if ERROR != nil {
		log.Fatalf("Error trying to build the database tables: %v", ERROR)
	}
}
