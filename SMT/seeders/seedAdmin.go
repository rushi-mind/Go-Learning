package seeders

import (
	"SMT/config"
	"SMT/models"
	"SMT/utility"
	"os"
)

func SeedFirstAdmin() {
	var admin models.Admin
	DB := config.DB
	DB.First(&admin)
	if admin.Id == 0 {
		admin.Id = 1
		admin.AdminId = "101425"
		admin.EmailId = "admin@gmail.com"
		admin.Password = utility.GetEncryptedPassword(os.Getenv("DEFAULT_ADMIN_PASSWORD"))
		DB.Create(&admin)
	}
}
