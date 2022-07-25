package main

import (
	"gorm-demo/config"
	"gorm-demo/models"
	"gorm-demo/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env => ", err)
	}
	config.InitConfig()
	models.SyncAllModels()

	routes.InitRoutes()

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), routes.Router))
}
