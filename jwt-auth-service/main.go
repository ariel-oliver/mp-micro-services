package main

import (
	"log"
	"net/http"
    
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/config"
	"github.com/ariel-oliver/mp-micro-services/jwt-auth-service/routers"
)

func main() {
	
	err := config.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	
	router := routers.InitRouter()

	
	log.Println("Starting server on :8083")
	err = http.ListenAndServe(":8083", router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
