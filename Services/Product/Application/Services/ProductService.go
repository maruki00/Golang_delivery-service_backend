package product_services

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	"errors"
)

type ProductService struct {
	productRepository product_domain_repositories.IProductRepository
}

func NewProductService(productRepository product_domain_repositories.IProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (obj *ProductService) Insert(dto *product_domain_dtos.InsertProductDTO) (bool, error) {

	res, err := obj.productRepository.Insert(&product_infrastructure_models.Product{
		Label: dto.Label,
		Type:  dto.Type,
		Price: dto.Price,
	})

	if err != nil {
		return false, errors.New("error")
	}

	if !res {
		return false, errors.New("could not create product reecord")
	}

	return true, nil
}
