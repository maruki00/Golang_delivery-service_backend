package order_domain_contracts

import (
	"context"
	order_domain_aggrigate "delivery/internal/order/Domain/Aggrigates"
	order_domain_entities "delivery/internal/order/Domain/Entities"
)

type IOrderRepository interface {
	Make(ctx context.Context, order *order_domain_aggrigate.OrderAggrigate) (*order_domain_aggrigate.OrderAggrigate, error)
	Cancel(ctx context.Context, id int) error
	Confirm(ctx context.Context, id int) error
	GetStatus(ctx context.Context, id int) (int, error)
	GetCustomerOrders(ctx context.Context, id int) ([]order_domain_entities.OrderEntity, error)
	GetByFingerPrint(ctx context.Context, fingerprint string) ([]order_domain_entities.OrderEntity, error)
}
