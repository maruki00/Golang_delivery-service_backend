package order_domain_contracts

import (
	"context"
	order_domain_aggrigate "delivery/Services/Order/Domain/Aggrigates"
	order_domain_entities "delivery/Services/Order/Domain/Entities"
)

type IOrderRepository interface {
	Make(context context.Context, order *order_domain_aggrigate.OrderAggrigate) (*order_domain_aggrigate.OrderAggrigate, error)
	Cancel(context context.Context, id int) error
	Confirm(context context.Context, id int) error
	GetStatus(context context.Context, id int) (int, error)
	GetCustomerOrders(context context.Context, id int) ([]order_domain_entities.OrderEntity, error)
	GetByFingerPrint(context context.Context, fingerprint string) ([]order_domain_entities.OrderEntity, error)
}
