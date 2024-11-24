package order_application_services

import (
	order_domain_aggrigate "delivery/Internal/Order/Domain/Aggrigates"
	order_domain_contracts "delivery/Internal/Order/Domain/Contracts"
	order_domain_dtos "delivery/Internal/Order/Domain/DTOs"
	order_domain_ports "delivery/Internal/Order/Domain/Ports"
	shared_domain_contracts "delivery/Internal/Shared/Domain/Contracts"
)

type OrderService struct {
	repository order_domain_contracts.IOrderRepository
	outputPort order_domain_ports.OrderOutputPort
}

func NewOrderService(
	repository order_domain_contracts.IOrderRepository,
	outputPort order_domain_ports.OrderOutputPort,
) *OrderService {
	return &OrderService{
		repository: repository,
		outputPort: outputPort,
	}
}

func (obj *OrderService) CreateOrder(dto order_domain_dtos.CreateNewOrderDTO) shared_domain_contracts.ViewModel {

	// orderAggrigate, err := order_factories.NewOrderAggrigate(dto.CostomerId, dto.Products)
	// if err != nil {
	// 	panic(err)
	// }

	// obj.repository.Make(orderAggrigate)
	return nil
}

func (obj *OrderService) Make(order *order_domain_aggrigate.OrderAggrigate) shared_domain_contracts.ViewModel {

	return nil
}

func (obj *OrderService) Cancel(id int) shared_domain_contracts.ViewModel {
	return nil
}

func (obj *OrderService) Confirm(id int) shared_domain_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetStatus(id int) shared_domain_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetCustomerOrders(id int) shared_domain_contracts.ViewModel {
	return nil
}

func (obj *OrderService) GetByFingerPrint(fingerprint string) shared_domain_contracts.ViewModel {
	return nil
}
