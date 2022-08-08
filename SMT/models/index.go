package models

import (
	"SMT/config"
)

func SyncModels() {
	var DB = config.DB
	DB.AutoMigrate(&Admin{})
	DB.AutoMigrate(&Department{})
	DB.AutoMigrate(&Student{})
	DB.AutoMigrate(&Faculty{})
	DB.AutoMigrate(&Attendace{})
	DB.AutoMigrate(&Assignment{})
	DB.AutoMigrate(&LeaveApplication{})
}
