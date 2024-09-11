package product_domain_ports

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_entities "delivery/Services/Product/Domian/Entities"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
)

type ProductInputPort interface {
	Insert(dto *product_domain_dtos.InsertProductDTO) (product_domain_entities.ProductEntity, error)
	Search(dto *product_domain_dtos.SearchProductDTO) ([]product_infrastructure_models.Product, error)
	Update(dto *product_domain_dtos.UpdateProductDTO) (product_domain_entities.ProductEntity, error)
	Delete(dto *product_domain_dtos.DeleteProductDTO) (product_domain_entities.ProductEntity, error)
	GetById(dto *product_domain_dtos.GetProductByIdDTO) (product_domain_entities.ProductEntity, error)
}
