package controllers

import (
	"github.com/capt4ce/ecommerce-microservice/products/models"
)

type (
	// For Get - /products
	ProductsResource struct {
		Data []models.Product `json:"data"`
	}
	// For Post/Put - /products
	ProductResource struct {
		Data models.Product `json:"data"`
	}
)
