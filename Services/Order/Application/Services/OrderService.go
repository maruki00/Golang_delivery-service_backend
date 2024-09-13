package main

import (
	"context"
	order_domain_contracts "delivery/Services/Order/Domain/Contracts"
	order_domain_dtos "delivery/Services/Order/Domain/DTOs"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
)

type OrderService struct {
	repository order_domain_contracts.IOrderRepository
	outputPort order_domain
}


func (obj *OrderService)CreateOrder(context context.Context, dto order_domain_dtos.CreateNewOrderDTO) shared_domain_contracts.ViewModel {
	
}