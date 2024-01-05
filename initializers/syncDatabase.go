package initializers

import "simple-rest-api-golang/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
