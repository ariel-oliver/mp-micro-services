package routers

import (
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/product-service/controllers"
)

func InitRouter() *http.ServeMux {
	//Inserir PUT
	router := http.NewServeMux()
	router.HandleFunc("POST /products", controllers.CreateProduct)
	router.HandleFunc("GET /products", controllers.GetProducts)
	router.HandleFunc("PUT /products/{id}", controllers.UpdateProduct)
	return router
}