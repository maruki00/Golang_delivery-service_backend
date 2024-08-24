package product_usergetway_controllers

import (
	product_services "delivery/Services/Product/Application/Services"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	product_Infrastructure_repository "delivery/Services/Product/Infrastructure/Repositories"

	"gorm.io/gorm"
)

type ProductController struct {
	service product_services.ProductService
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		service: &product_services.NewProductService(
			product_Infrastructure_repository.NewProductRepository(db, &product_infrastructure_models.Product{}),
		),
	}
}
