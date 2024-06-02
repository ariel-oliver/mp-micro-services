package routers

import (
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
