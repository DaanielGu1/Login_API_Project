package main

import (
	"fmt"
	"log"

	"github.com/DaanielGu1/Login_API_Project/database"
)

func main() {
	db := database.CreateDatabase()
	defer db.Close()

	id_user, err := database.CreateNewUser(db, "Daniel", "Guimaraes", "11/11/2011", "senha123")
	if err != nil {
		log.Fatalf("Error trying to create a new user: %v", err)
	}
	
	id_group, err := database.CreateNewGroup(db, "New_Group")
	if err != nil {
		log.Fatalf("Error trying to create a new group: %v", err)
	}

	fmt.Printf("Success when creating user! ID: %d\n", id_user)
	fmt.Printf("Success when creating group! ID: %d\n", id_group)
}
