package product_domain_ports

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
)

type ProductInputPort interface {
	Insert(dto *product_domain_dtos.InsertProductDTO) shared_domain_contracts.ViewModel
	Search(dto *product_domain_dtos.SearchProductDTO) shared_domain_contracts.ViewModel
	Update(dto *product_domain_dtos.UpdateProductDTO) shared_domain_contracts.ViewModel
	Delete(dto *product_domain_dtos.DeleteProductDTO) shared_domain_contracts.ViewModel
	GetById(dto *product_domain_dtos.GetProductByIdDTO) shared_domain_contracts.ViewModel
}
