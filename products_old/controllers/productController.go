package controllers

import(
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

// to get the product list
func GetProducts(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	collection := context.DbCollection("products")
	repo := mgo.Collection{collection}
	products :=repo.GetAll()
	j, err := json.Marshal(ProductsResource{Data: products})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}


// adding a new product to the database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	
}


// getting a specific product with corresponding Id
func GetProductById(w http.ResponseWriter, r *http.Request) {
	
}


// delete a product from the database
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	
}