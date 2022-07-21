package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-demo/config"
	"rest-demo/structs"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	getUsersResult, err := config.DB.Query(`select * from users`)
	if err != nil {
		log.Fatal(err)
	}

	var users []structs.User
	for getUsersResult.Next() {
		var user structs.User
		_ = getUsersResult.Scan(&user.Id, &user.Name, &user.Pin)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	var user structs.User
	config.DB.QueryRow(`select * from users where id = ?`, id).Scan(&user.Id, &user.Name, &user.Pin)
	json.NewEncoder(w).Encode(struct {
		Status int           `json:"status"`
		Data   *structs.User `json:"data"`
	}{1, &user})
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	res, _ := config.DB.Exec(`insert into users(name, pin) values(?, ?)`, user.Name, user.Pin)
	temp, _ := res.LastInsertId()
	responseMessage := "User added successfully. Last inserted id: " + strconv.Itoa(int(temp))
	json.NewEncoder(w).Encode(struct {
		Status  int
		Message string
	}{1, responseMessage})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user structs.User
	id := mux.Vars(r)["id"]
	err := config.DB.QueryRow(`select * from users where id = ?`, id).Scan(&user.Id, &user.Name, &user.Pin)
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Status  int
			Message string
		}{0, "Invalid user ID"})
		fmt.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Exec(`update users set name = ?, pin = ? where id = ?`, user.Name, user.Pin, user.Id)
	json.NewEncoder(w).Encode(struct {
		Status  int
		Message string
	}{1, "User Updated Successfully"})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	tempId := 0
	config.DB.QueryRow(`select id from users where id = ?`, id).Scan(&tempId)
	if tempId == 0 {
		json.NewEncoder(w).Encode(struct {
			Status  int
			Message string
		}{0, "Invalid User-ID"})
		return
	}
	_, err := config.DB.Exec(`delete from users where id = ?`, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(struct {
		Status  int
		Message string
	}{1, "User deleted successfully"})
}
