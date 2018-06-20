package routers

import (
	"github.com/capt4ce/ecommerce-microservice/products/controllers"
	"github.com/gorilla/mux"
)

func setMovieRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/products", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.DeleteMovie).Methods("DELETE")
	return router
}
