package routers

import (
	"github.com/ariel-oliver/mp-micro-services/product-service/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	//Inserir PUT
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	return router
}
