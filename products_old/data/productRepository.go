package data

import (
	"github.com/capt4ce/ecommerce-microservice/products/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


func (r *mgo.Collection) Create(product *models.Product) error {
	obj_id := bson.NewObjectId()
	product.Id = obj_id
	err := r.Insert(&product)
	return err
}

func (r *mgo.Collection) GetAll() []models.Product {
	var products []models.Product
	iter := r.Find(nil).Iter()
	result := models.Product{}
	for iter.Next(&result) {
		products = append(products, result)
	}
	return products
}

func (r *mgo.Collection) GetById(id string) (product models.Product, err error) {
	err = r.FindId(bson.ObjectIdHex(id)).One(&product)
	return err
}

func (r *mgo.Collection) Delete(id string) error {
	err := r.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
