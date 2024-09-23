package order_application_services

import (
	order_domain_aggrigate "delivery/Services/Order/Domain/Aggrigates"
	order_domain_contracts "delivery/Services/Order/Domain/Contracts"
	order_domain_dtos "delivery/Services/Order/Domain/DTOs"
	order_domain_ports "delivery/Services/Order/Domain/Ports"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
	"net/http"
)

type OrderService struct {
	repository order_domain_contracts.IOrderRepository
	outputPort order_domain_ports.OrderOutputPort
	// ctx        context.Context
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

	price := float32(0)
	var client *http.Client

	client.Post("http://locahost:3000/api/product/get", "application/json", map[string]string{})

}

func (obj *OrderService) Make(order *order_domain_aggrigate.OrderAggrigate) shared_domain_contracts.ViewModel {

}

func (obj *OrderService) Cancel(id int) shared_domain_contracts.ViewModel {

}

func (obj *OrderService) Confirm(id int) shared_domain_contracts.ViewModel {

}

func (obj *OrderService) GetStatus(id int) shared_domain_contracts.ViewModel {

}

func (obj *OrderService) GetCustomerOrders(id int) shared_domain_contracts.ViewModel {

}

func (obj *OrderService) GetByFingerPrint(fingerprint string) shared_domain_contracts.ViewModel {

}
