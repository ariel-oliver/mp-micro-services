package main

import (
	"log"
	"net/http"

	"github.com/ariel-oliver/mp-micro-services/product-service/routers"
)

func main() {
	router := routers.InitRouter()
	log.Fatal(http.ListenAndServe(":8081", router))
}
