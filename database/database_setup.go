package database

import (
	"database/sql"
	"log"
	"os"

	_"github.com/mattn/go-sqlite3"
	_"go.mongodb.org/mongo-driver/v2/x/mongo/driver/description"
	_"golang.org/x/text/date"
)

func CreateDatabase(db *sql.DB) {

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
}

/* func CreateNewUser(db *sql.DB, user User) int {
	query := `INSERT INTO users VALUES($1, $2, $3, $4, $5, %6) RETURNING id`

	var id int
	
	err := db.QueryRow(query, user.Username, user.First_name, user.Last_name, user.Birthday, user.Email, user.Password).Scan(&id)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return id
}

func CreateNewGroup(db *sql.DB, group Group) int {
	query := `INSERT INTO groups VALUES($1, $2) RETURNING id`

	var id int

	err := db.QueryRow(query, group.Name, group.Description).Scan(&id)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return id
}

*/