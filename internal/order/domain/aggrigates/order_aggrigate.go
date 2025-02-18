package order_domain_aggrigate

import (
	order_entities "delivery/internal/order/Domain/Entities"
	product_infra_models "delivery/internal/product/infra/Models"
)

type OrderAggrigate struct {
	Order order_entities.OrderEntity
	Items []*product_infra_models.Product
	Price float32
}
