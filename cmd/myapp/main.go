package main

import (
	"encoding/json"
	"fmt"
	"go-start/pkg/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products := db.GetProducts(db.Database)

		for index, product := range products {
			fmt.Println(index, product)
		}

		productsJson, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "Error encoding products", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productsJson)
	case http.MethodPost:
		product := db.Product{}
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			fmt.Println("Error decoded body", err)
			http.Error(w, "Error decoded body", http.StatusInternalServerError)
			return
		}

		db.InsertProduct(db.Database, &product)

		productsJson, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Error encoding products", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(productsJson)
	case http.MethodDelete:
		productDelete := db.Product{}

		err := json.NewDecoder(r.Body).Decode(&productDelete)
		if err != nil {
			fmt.Println("Error decoded body: ", err)
			http.Error(w, "Error decoded body", http.StatusInternalServerError)
			return
		}

		err = db.DeleteProductById(db.Database, productDelete.Id)
		if err != nil {
			fmt.Println("Delete product error: ", err)
			http.Error(w, "Delete product error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not found", http.StatusInternalServerError)
	}
}

func main() {
	// Load .env
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file", errEnv)
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Println(dbUser, dbName, dbPassword, dbSslMode)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s  sslmode=%s ", dbUser, dbPassword, dbName, dbSslMode)

	db.InitDb(connStr)

	// Server
	http.HandleFunc("/products", productsHandler)

	addr := fmt.Sprintf("%s:%s", host, port)
	http.ListenAndServe(addr, nil)
}
