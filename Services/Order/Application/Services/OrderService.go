package main

import order_domain_contracts "delivery/Services/Order/Domain/Contracts"

type OrderService struct {
	repository order_domain_contracts.IOrderRepository
	outputPort order_domain
}
