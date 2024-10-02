package order_domain_aggrigate

import (
	order_entities "delivery/Services/Order/Domain/Entities"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
)

type OrderAggrigate struct {
	Order order_entities.OrderEntity
	Items []*product_infrastructure_models.Product
	Price float32
}
