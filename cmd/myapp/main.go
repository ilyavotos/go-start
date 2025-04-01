package main

import (
	"fmt"
	"go-start/internal/products"
	"go-start/pkg/db"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

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

	// Routing
	http.HandleFunc("/products", products.Handler)
	http.HandleFunc("/products/{id}", products.HandlerId)

	// Server
	addr := fmt.Sprintf("%s:%s", host, port)
	http.ListenAndServe(addr, nil)

}
