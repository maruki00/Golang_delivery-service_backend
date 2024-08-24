package product_usergetway_controllers

import (
	product_services "delivery/Services/Product/Application/Services"
	product_Infrastructure_repository "delivery/Services/Product/Infrastructure/Repositories"
)

type ProductController struct {
	service product_services.ProductService
}

func NewProductController() *ProductController {
	repo := product_Infrastructure_repository.NewProductRepository()
	return &ProductController{
		service: &product_services.NewProductService(),
	}
}
