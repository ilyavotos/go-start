package products

import (
	"encoding/json"
	"fmt"
	"go-start/pkg/db"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	productRepository := NewSQLProductRepository(db.Database)
	productService := NewRepoProductService(productRepository)
	switch r.Method {
	case http.MethodGet:
		products, _ := productService.GetAllProduct()

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
		product := Product{}
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			fmt.Println("Error decoded body", err)
			http.Error(w, "Error decoded body", http.StatusInternalServerError)
			return
		}

		productService.CreateProduct(&product)

		productsJson, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Error encoding products", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(productsJson)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func HandlerId(w http.ResponseWriter, r *http.Request) {
	productRepository := NewSQLProductRepository(db.Database)
	productService := NewRepoProductService(productRepository)

	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Id error", http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		product, err := productService.GetProduct(idInt)

		if err != nil {
			http.Error(w, "GetProductById error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		productJson, err := json.Marshal(product)
		if err != nil {
			http.Error(w, "Error encoding products", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(productJson)
	case http.MethodDelete:
		err = productService.DeleteProduct(idInt)
		if err != nil {
			fmt.Println("Delete product error: ", err)
			http.Error(w, "Delete product error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
