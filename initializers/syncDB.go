package initializers

import (
	"log"
	"social-journal/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Journal{})

	if err != nil {
		log.Fatal("Failed to migrate model", err)
	}
}
