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

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product
	config.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}
	config.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
	var product models.Product
	product.ID = id
	temp := config.DB.First(&product)
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
	json.NewDecoder(r.Body).Decode(&product)
	config.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
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
	var product models.Product
	product.ID = id
	deleteResult := config.DB.Delete(&product)
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
	}{1, "Product deleted successfully"})
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
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
	var product models.Product
	product.ID = id
	findResult := config.DB.Find(&product)
	var count int64
	findResult.Count(&count)
	if count == 0 {
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{0, "Invalid ID entered"})
		return
	}
	json.NewEncoder(w).Encode(product)
}
