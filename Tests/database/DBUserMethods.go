package database

import "Treino-em-go/models"

func GetUserByID(db *DB, id int) *models.User {
	for _, user := range db.users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func GetUserByName(db *DB, name string) *models.User {
	for _, user := range db.users {
		if user.Name == name {
			return &user
		}
	}

	return nil
}