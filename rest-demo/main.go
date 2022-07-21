package main

import (
	"log"
	"net/http"
	"os"
	"rest-demo/config"
	"rest-demo/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env not found")
	}

	config.Init()
	defer config.DB.Close()

	routes.Init()

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), routes.Router))
}
