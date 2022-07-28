package main

import (
	"SMT/config"
	"SMT/models"
	"SMT/routes"
	"SMT/seeders"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// loading .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not found")
	}

	// DB connection
	config.InitDBConfig(os.Getenv("DB_CONNECTION_STRING"))

	// syncing models
	models.SyncModels()

	// seeding minimum required data in DB
	seeders.SeedFirstAdmin()

	// initializing routes
	routes.InitRoutes()
	log.Fatal(routes.Router.Run(os.Getenv("PORT")))

}
