package models

import (
	"SMT/config"
)

func SyncModels() {
	var DB = config.DB
	_ = DB.AutoMigrate(&Admin{})
	_ = DB.AutoMigrate(&Department{})
	_ = DB.AutoMigrate(&Student{})
	// DB.AutoMigrate((&Faculty{}))
}
