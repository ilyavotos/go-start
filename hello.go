package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Product struct {
	id      int
	model   string
	company string
	price   int
}

func getProductById(db *sql.DB, id int) Product {
	row := db.QueryRow("select * from products where id = $1", id)

	product := Product{}

	err := row.Scan(&product.id, &product.model, &product.company, &product.price)
	if err != nil {
		panic(err)
	}
	return product
}

func getProducts(db *sql.DB) []Product {
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.id, &product.model, &product.company, &product.price)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func insertProduct(db *sql.DB) int {
	var id int
	db.QueryRow("insert into products (model, company, price) values ($1,$2,$3) returning id", "Android", "Xiaomi", 1500).Scan(&id)
	return id
}

func deleteProductById(db *sql.DB, id int) {
	result, err := db.Exec("delete from products where id = $1", id)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Count deleted products:", count)
}

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	fmt.Println(dbUser, dbName, dbPassword, dbSslMode)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s  sslmode=%s ", dbUser, dbPassword, dbName, dbSslMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// productId := insertProduct(db)
	// fmt.Println(productId)

	deleteProductById(db, 6)

	products := getProducts(db)

	for index, product := range products {
		fmt.Println(index, product)
	}

	product1 := getProductById(db, 1)
	fmt.Println(product1)

}
