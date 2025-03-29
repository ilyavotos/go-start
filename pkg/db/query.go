package db

import (
	"database/sql"
	"fmt"
)

type Product struct {
	id      int
	model   string
	company string
	price   int
}

func GetProductById(db *sql.DB, id int) Product {
	row := db.QueryRow("select * from products where id = $1", id)

	product := Product{}

	err := row.Scan(&product.id, &product.model, &product.company, &product.price)
	if err != nil {
		panic(err)
	}
	return product
}

func GetProducts(db *sql.DB) []Product {
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

func InsertProduct(db *sql.DB) int {
	var id int
	db.QueryRow("insert into products (model, company, price) values ($1,$2,$3) returning id", "Android", "Xiaomi", 1500).Scan(&id)
	return id
}

func DeleteProductById(db *sql.DB, id int) {
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
