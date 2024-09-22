package order_factories

import order_application_services "delivery/Services/Order/Application/Services"

func NewOrderService() *order_application_services.OrderService {
	return &order_application_services.OrderService{}
}
