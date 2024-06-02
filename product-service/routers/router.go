package routers

import (
	"github.com/ariel-oliver/mp-micro-services/product-service/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	return router
}
