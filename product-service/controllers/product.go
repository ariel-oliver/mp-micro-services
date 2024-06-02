package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/product-service/config"
	"github.com/ariel-oliver/mp-micro-services/product-service/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	stmt, err := config.DB.Prepare("INSERT INTO products (name, price) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(product.Name, product.Price)
	if err != nil {
		http.Error(w, "Could not create product", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			http.Error(w, "Error scanning product", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    var updatedProduct models.Product
    err := json.NewDecoder(r.Body).Decode(&updatedProduct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }
    updatedProduct.ID = id

    err = models.UpdateProduct(updatedProduct)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedProduct)
}