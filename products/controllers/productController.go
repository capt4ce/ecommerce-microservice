package controllers

import(
	"github.com/capt4ce/ecommerce-microservice/products/common"
	"github.com/capt4ce/ecommerce-microservice/products/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	mongo_session = common.GetDbSession()
	defer mongo_session.Close()

	var products []*models.Product
	if err := mongo_session.DB("products").C("product").
		Find(nil).Sort("-when").Limit(100).All(&products); err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write it out
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func GetProductById(w http.ResponseWriter, r *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
