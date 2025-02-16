package product_domain_ports

import (
	"context"
	product_domain_dtos "delivery/internal/product/Domian/DTOS"
	shared_domain_contracts "delivery/internal/shared/Domain/Contracts"
)

type ProductInputPort interface {
	Insert(ctx context.Context, dto *product_domain_dtos.InsertProductDTO) shared_domain_contracts.ViewModel
	Search(ctx context.Context, dto *product_domain_dtos.SearchProductDTO) shared_domain_contracts.ViewModel
	Update(ctx context.Context, dto *product_domain_dtos.UpdateProductDTO) shared_domain_contracts.ViewModel
	Delete(ctx context.Context, dto *product_domain_dtos.DeleteProductDTO) shared_domain_contracts.ViewModel
	GetById(ctx context.Context, dto *product_domain_dtos.GetProductByIdDTO) shared_domain_contracts.ViewModel
	MultipleProducts(ctx context.Context, dto *product_domain_dtos.MultipleProductsDTO) shared_domain_contracts.ViewModel
}
