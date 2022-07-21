package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func stringSort(temp string) string {
	s := strings.Split(temp, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	// var inputChoice int
	// fmt.Printf("\t%s, Enter any number. Available choices: %v\t=>\t", player, available)
	// fmt.Print("Enter your choice: ")
	// _, err := fmt.Scan(&inputChoice)
	// if err == nil {
	// 	fmt.Printf("your choice %v %T\n", inputChoice, inputChoice)
	// } else {
	// 	fmt.Printf("error: %v %T\n", err, err)
	// }

	// var temp string = stringSort("bcda")
	// fmt.Println(temp[0])

	// var temp = []int{}

	// temp = append(temp, 1)
	// temp = append(temp, 2, 3, 4, 5)

	// temp = append(temp[:5])

	// fmt.Printf("%v : %T\n", temp, temp)
	// fmt.Println(len(temp), ":", cap(temp))
	// temp = append(temp, 6)
	// fmt.Printf("%v : %T\n", temp, temp)
	// fmt.Println(len(temp), ":", cap(temp))
	// temp = append(temp, 7, 8, 9, 10, 11)
	// fmt.Printf("%v : %T\n", temp, temp)
	// fmt.Println(len(temp), ":", cap(temp))
	// temp = append(temp, 12, 13)
	// fmt.Printf("%v : %T\n", temp, temp)
	// fmt.Println(len(temp), ":", cap(temp))
	// temp = append(temp, 6, 7, 8, 9, 10, 11, 12, 13)
	// temp[3] = 4
	// temp[4] = 5
	// temp[5] = 6

	// temp := make(map[string]string)
	// fmt.Printf("%v : %T\n", temp, temp)
	// fmt.Println()

	// str := "this"
	// fmt.Println(string(str[1]))
	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":null}`

	var result map[string]interface{}
	json.Unmarshal([]byte(birdJson), &result)
	type bird struct {
		pigeon string
		eagle  string
	}
	var temp = result["birds"].(map[string]interface{})
	for key, value := range temp {
		fmt.Println("key:", key, " - value:", value)
	}
	fmt.Println(temp)
	fmt.Println(result)
}
