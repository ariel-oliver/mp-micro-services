package main

import (
	"log"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/order-service/config"
	"github.com/ariel-oliver/mp-micro-services/order-service/routers"
)

func main() {
	config.InitDB()
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8082", router))
}
