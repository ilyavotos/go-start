package products

import (
	"database/sql"
	"errors"
)

type ProductRepository interface {
	FindById(id int) (Product, error)
	FindAll() ([]Product, error)
	Create(product *Product) error
	DeleteById(id int) error
}

type SQLProductRepository struct {
	db *sql.DB
}

func NewSQLProductRepository(db *sql.DB) *SQLProductRepository {
	return &SQLProductRepository{db}
}

func (repo *SQLProductRepository) FindById(id int) (Product, error) {
	row := repo.db.QueryRow("select * from products where id = $1", id)

	product := Product{}

	err := row.Scan(&product.Id, &product.Model, &product.Company, &product.Price)

	return product, err
}

func (repo *SQLProductRepository) FindAll() ([]Product, error) {
	products := []Product{}

	rows, err := repo.db.Query("select * from products")
	if err != nil {
		return products, err
	}

	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.Id, &product.Model, &product.Company, &product.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}

	return products, err
}

func (repo *SQLProductRepository) Create(product *Product) error {
	err := repo.db.QueryRow("insert into products (model, company, price) values ($1,$2,$3) returning id", product.Model, product.Company, product.Price).Scan(&product.Id)
	return err
}

func (repo *SQLProductRepository) DeleteById(id int) error {
	result, err := repo.db.Exec("delete from products where id = $1", id)
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
