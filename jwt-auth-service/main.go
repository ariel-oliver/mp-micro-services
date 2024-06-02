package main

import (
	"log"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/config"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/routers"
)

func main() {
	config.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8083", router))
}
