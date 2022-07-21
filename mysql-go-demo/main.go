package main

import (
	"fmt"
	"log"
	"mysql-go-demo/config"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("let the game begin")

	err := godotenv.Load(".env")
	if err != nil {
		panic(".env not found")
	}

	DB := config.Init()
	defer DB.Close()

	var effectiveHours string
	DB.QueryRow(`select effective_hours from attendance where sr = 106`).Scan(&effectiveHours)
	if len(effectiveHours) > 0 {
		fmt.Println("Effective hours: ", effectiveHours)
	}
	fmt.Println()

	// var err error
	selectResult, err := DB.Query(`select * from attendance where id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	temp := make([]string, 7)
	fmt.Println("All attendance of user-1: ")
	for selectResult.Next() {
		err = selectResult.Scan(&temp[0], &temp[1], &temp[2], &temp[3], &temp[4], &temp[5], &temp[6])
		if err != nil {
			log.Fatal(err)
		}
		for _, e := range temp {
			fmt.Print(e, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	insertResult, err := DB.Exec(`insert into users(name, pin) values("TempUser", 4444)`)
	if err != nil {
		log.Fatal(err)
	}
	lastInsertedId, err := insertResult.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Last inserted id: ", lastInsertedId)
	fmt.Println()

	deleteResult, err := DB.Exec(`delete from users where id in (?)`, lastInsertedId)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err := deleteResult.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of rows deleted: ", affectedRows)
	fmt.Println()

	var totalNumRows int
	DB.QueryRow(`select count(id) from users as total`).Scan(&totalNumRows)
	fmt.Println("Total Number of Rows: ", totalNumRows)
	fmt.Println()

	logsCount, err := DB.Query(`select date, count(id) from attendance where id = 1 group by date`)
	if err != nil {
		panic(err)
	}
	for logsCount.Next() {
		var date string
		var count int
		_ = logsCount.Scan(&date, &count)
		fmt.Println("Date: ", date, "Total Entries: ", count)
	}
}
