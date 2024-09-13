package order_domain_contracts

import (
	order_domain_aggrigate "delivery/Services/Order/Domain/Aggrigates"
	order_domain_entities "delivery/Services/Order/Domain/Entities"
)

type IOrderRepository interface {
	Make(order *order_domain_aggrigate.OrderAggrigate) (*order_domain_aggrigate.OrderAggrigate, error)
	Cancel(id int) error
	Confirm(id int) error
	GetStatus(id int) (int, error)
	GetCustomerOrders(id int) ([]order_domain_entities.OrderEntity, error)
	GetByFingerPrint(fingerprint string) ([]order_domain_entities.OrderEntity, error)
}
