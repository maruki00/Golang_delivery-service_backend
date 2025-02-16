package order_domain_aggrigate

import (
	order_entities "delivery/internal/order/Domain/Entities"
	product_infrastructure_models "delivery/internal/product/Infrastructure/Models"
)

type OrderAggrigate struct {
	Order order_entities.OrderEntity
	Items []*product_infrastructure_models.Product
	Price float32
}
