package models

import (
	"SMT/config"
)

func SyncModels() {
	var DB = config.DB
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Department{})
	DB.AutoMigrate(&Student{})
	DB.AutoMigrate((&Faculty{}))
}
