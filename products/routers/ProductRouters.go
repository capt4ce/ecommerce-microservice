package routers

import (
	"github.com/gorilla/mux"
	"github.com/capt4ce/ecommerce-microservice/products/controllers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// Routes for the User entity
	router = SetUsersRouters(router)
	return router
}

func setProductRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	return router
}
