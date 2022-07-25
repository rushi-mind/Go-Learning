package models

import "gorm-demo/config"

func SyncAllModels() {
	config.DB.AutoMigrate(&Product{})
	config.DB.AutoMigrate(&Shop{})
}
