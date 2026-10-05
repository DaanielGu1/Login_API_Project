package database

import (
	"Treino-em-go/models"
	"slices"
)

func GroupExists(db *DB, group models.Group) bool {
	if slices.Contains(db.groups, group) {
		return true
	}

	return false
}