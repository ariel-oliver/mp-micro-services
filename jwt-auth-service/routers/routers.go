package routers

import (
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/controllers"
)

func InitRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("POST /register", controllers.CreateUser)
	router.HandleFunc("POST /login", controllers.Login)
	router.HandleFunc("PUT /user/{id}", controllers.UpdateUser)
	router.HandleFunc("DELETE /user/{id}", controllers.DeleteUser)
	return router
}
