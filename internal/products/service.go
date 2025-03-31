package products

import (
	"database/sql"
	"errors"
)

type Product struct {
	Id      int    `json:"id,omitempty"`
	Model   string `json:"model"`
	Company string `json:"company"`
	Price   int    `json:"price"`
}

func GetProductById(db *sql.DB, id int) (Product, error) {
	row := db.QueryRow("select * from products where id = $1", id)

	product := Product{}

	err := row.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProducts(db *sql.DB) []Product {
	rows, err := db.Query("select * from products")
	if err != nil {
		panic(err)
	}
	products := []Product{}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products
}

func InsertProduct(db *sql.DB, product *Product) {
	db.QueryRow("insert into products (model, company, price) values ($1,$2,$3) returning id", product.Model, product.Company, product.Price).Scan(&product.Id)
}

func DeleteProductById(db *sql.DB, id int) error {
	result, err := db.Exec("delete from products where id = $1", id)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Product id not found")
	}
	return nil
}
