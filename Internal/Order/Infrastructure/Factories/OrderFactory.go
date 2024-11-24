package order_factories

import order_application_services "delivery/Internal/Order/Application/Services"

func NewOrderService() *order_application_services.OrderService {
	return &order_application_services.OrderService{}
}
