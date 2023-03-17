package initializers

import "github.com/ndy-corp/1.src/midterm-1/src-code-new/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
