package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(uuid.New())

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/products", getProducts).Methods("GET")

	// db := connect()
	// defer db.Close()

	// product := Product{ID: uuid.New(), Name: "My Product", Quantity: 10, Price: 49.99}
	// fmt.Println(product)

	// Starting Server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Get Products
// Get Products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get Connect
	db := connect()
	defer db.Close()

	// Creating Products Slice
	var products []Product
	if err := db.Model(&products).Select(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Returning Products
	json.NewEncoder(w).Encode(products)
}
