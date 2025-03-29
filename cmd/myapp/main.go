package main

import (
	"fmt"
	"go-start/pkg/db"
	"log"
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

	fmt.Println(dbUser, dbName, dbPassword, dbSslMode)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s  sslmode=%s ", dbUser, dbPassword, dbName, dbSslMode)

	errDb := db.InitDb(connStr)
	if errDb != nil {
		log.Fatal("Db error", errDb)
	}

	db.DeleteProductById(db.Database, 6)

	products := db.GetProducts(db.Database)

	for index, product := range products {
		fmt.Println(index, product)
	}

	product1 := db.GetProductById(db.Database, 1)
	fmt.Println(product1)

}
