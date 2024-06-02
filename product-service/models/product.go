package models

import (
    "errors"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}
func UpdateProduct(updatedProduct Product) error {
    for i, p := range products {
        if p.ID == updatedProduct.ID {
            products[i] = updatedProduct
            return nil
        }
    }
    return errors.New("Product not found")
}
