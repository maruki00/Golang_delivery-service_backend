package product_services

import (
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
)

type ProductService struct {
	productRepository *product_domain_repositories.IProductRepository
}

func NewProductService(productRepository product_domain_repositories.IProductRepository) *ProductService {
	return &ProductService{
		productRepository: &productRepository,
	}
}
