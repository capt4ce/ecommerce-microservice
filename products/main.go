package main

import (
	"log"
	"net/http"

	"github.com/capt4ce/ecommerce-microservice/products/common"
	"github.com/capt4ce/ecommerce-microservice/products/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening [products]...")
	server.ListenAndServe()
}