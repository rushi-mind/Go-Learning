package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Layman struct {
	RocK string `json:"temp1"`
	Rock string `json:"temp2"`
}

var data []Layman

func post(w http.ResponseWriter, r *http.Request) {
	var j Layman
	json.NewDecoder(r.Body).Decode(&j)
	fmt.Println(j)
	data = append(data, j)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/api/post", post).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", r))
}
