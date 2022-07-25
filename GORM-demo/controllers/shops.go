package controllers

import (
	"encoding/json"
	"fmt"
	"gorm-demo/config"
	"gorm-demo/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetShops(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shops []models.Shop
	config.DB.Find(&shops)
	json.NewEncoder(w).Encode(shops)
}

func AddShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var shop models.Shop
	err := json.NewDecoder(r.Body).Decode(&shop)
	if err != nil {
		log.Fatal(err)
	}
	config.DB.Create(&shop)
	json.NewEncoder(w).Encode(shop)
}

func UpdateShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tempId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{0, "Invalid ID entered"})
		return
	}
	id := uint(tempId)
	var shop models.Shop
	shop.ID = id
	temp := config.DB.First(&shop)
	var count int64
	temp.Count(&count)
	if count == 0 {
		err := json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{0, "Invalid ID entered"})
		fmt.Println(err)
		return
	}
	json.NewDecoder(r.Body).Decode(&shop)
	config.DB.Save(&shop)
	json.NewEncoder(w).Encode(shop)
}

func DeleteShop(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tempId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{0, "Invalid ID entered"})
		return
	}
	id := uint(tempId)
	var shop models.Shop
	shop.ID = id
	deleteResult := config.DB.Delete(&shop)
	if deleteResult.RowsAffected == 0 {
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{0, "Invalid ID entered"})
		return
	}
	json.NewEncoder(w).Encode(struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{1, "Shop deleted successfully"})
}
