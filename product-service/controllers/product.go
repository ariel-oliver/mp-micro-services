package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Product struct {
	Uuid        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
}

type Products struct {
	Products []Product
}

func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println("erro: ", err.Error())
	}
	defer jsonFile.Close()

	data, _ := io.ReadAll(jsonFile)
	return data
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	data := loadData()

	var products Products
	json.Unmarshal(data, &products)

	for _, v := range products.Products {
		if v.Uuid == productID {
			product, _ := json.Marshal(v)
			w.Write([]byte(product))
		}
	}
}
