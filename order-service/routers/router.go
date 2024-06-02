package routers

import (
	"github.com/ariel-oliver/mp-micro-services/order-service/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	return router
}
