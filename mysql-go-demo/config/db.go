package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	DB, connectionError := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp(127.0.0.1:3306)/"+os.Getenv("DATABASE"))
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	return DB
}
