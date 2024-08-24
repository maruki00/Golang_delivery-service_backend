package product_services

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
)

type ProductService struct {
	productRepository *product_domain_repositories.IProductRepository
	// outputPort        *product_domain_ports.ProductOutputPort
}

func NewProductService(productRepository product_domain_repositories.IProductRepository,
	//outputPort product_domain_ports.NewProductOutputPort,
	model interface{}) *ProductService {
	return &ProductService{
		productRepository: &productRepository,
		// outputPort:        &outputPort,
	}
}

func (this ProductService) AddProduct(dto product_domain_dtos.ProductDTO) (error, ViewModel) {
	this.outputPortl.
}