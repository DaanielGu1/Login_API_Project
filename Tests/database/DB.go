package database

import "Treino-em-go/models"

type DB struct {
	users  []models.User
	groups []models.Group
}

var Database = &DB{}
