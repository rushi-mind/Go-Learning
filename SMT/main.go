package main

import (
	"SMT/config"
	"SMT/models"
	"SMT/routes"
	"SMT/seeders"
	"SMT/validations"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// loading .env
	fmt.Println("loading .env")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not found")
	}

	// connecting to DB
	fmt.Println("connecting to DB")
	config.InitDBConfig(os.Getenv("DB_CONNECTION_STRING"))

	// syncing models
	fmt.Println("syncing models")
	models.SyncModels()

	// registering validations
	fmt.Println("registering validations")
	validations.RegisterValidations()

	// seeding minimum required data in DB
	fmt.Println("seeding minimum required data in DB")
	seeders.SeedAll()

	// initializing routes
	fmt.Println("initializing routes")
	routes.InitRoutes()

	// starting server
	fmt.Println("Server setup successful")
	log.Fatal(routes.Router.Run(os.Getenv("PORT")))
}
