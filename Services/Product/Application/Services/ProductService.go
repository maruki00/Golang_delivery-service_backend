package product_services

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_entities "delivery/Services/Product/Domian/Entities"
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	"errors"
	"fmt"
)

type ProductService struct {
	productRepository product_domain_repositories.IProductRepository
}

func NewProductService(productRepository product_domain_repositories.IProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (obj *ProductService) Insert(dto *product_domain_dtos.InsertProductDTO) (product_domain_entities.ProductEntity, error) {

	res, err := obj.productRepository.Insert(&product_infrastructure_models.Product{
		Label: dto.Label,
		Type:  dto.Type,
		Price: dto.Price,
	})

	if err != nil {
		return nil, errors.New("error")
	}

	return res, nil
}

func (obj *ProductService) Search(dto *product_domain_dtos.SearchProductDTO) ([]product_domain_entities.ProductEntity, error) {

	res, err := obj.productRepository.Search(dto.GetQuery())

	if err != nil {
		return nil, errors.New("error")
	}

	return res, nil
}

func (obj *ProductService) Update(dto *product_domain_dtos.UpdateProductDTO) (product_domain_entities.ProductEntity, error) {

	res, err := obj.productRepository.Update(dto.Id, map[string]interface{}{
		"label": dto.Label,
		"type":  dto.Type,
		"price": dto.Price,
	})

	if err != nil {
		return nil, errors.New("error")
	}

	return res, nil
}

func (obj *ProductService) Delete(dto *product_domain_dtos.DeleteProductDTO) (product_domain_entities.ProductEntity, error) {

	res, err := obj.productRepository.Delete(dto.Id)

	if err != nil {
		return nil, fmt.Errorf("could not delete the product : %s", err.Error())
	}

	return res, nil
}

func (obj *ProductService) GetById(dto *product_domain_dtos.GetProductByIdDTO) (product_domain_entities.ProductEntity, error) {

	res, err := obj.productRepository.GetById(dto.Id)

	if err != nil {
		return nil, errors.New("could not get the product or doesnt exsits")
	}
	return res, nil
}
