package models

import (
	"SMT/config"
	"log"
)

func SyncModels() {
	var DB = config.DB
	var err error
	err = DB.AutoMigrate(&Admin{})
	if err != nil {
		log.Fatal("Failed to migrate models")
	}
	err = DB.AutoMigrate(&Department{})
	if err != nil {
		log.Fatal("Failed to migrate models")
	}
	// DB.AutoMigrate(&Student{})
	// DB.AutoMigrate((&Faculty{}))
}
