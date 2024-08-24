package orderaggrigate

import (
	order_entities "delivery/Services/Order/Domain/Entities"
	product_domain_entities "delivery/Services/Product/Domian/Entities"
)

type OrderAggrigate struct {
	orderEntity order_entities.OrderEntity
	items       []product_domain_entities.ProductEntity
}
