package routers

import (
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/product-service/controllers"
)

func InitRouter() *http.ServeMux {
	//Inserir PUT
	router := http.NewServeMux()
	router.HandleFunc("GET /products", controllers.ListProducts)
	router.HandleFunc("GET /product/{id}", controllers.GetProductById)
	return router
}
