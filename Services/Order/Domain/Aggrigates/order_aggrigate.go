package order_domain_aggrigate

import (
	order_entities "delivery/Services/Order/Domain/Entities"
	product_domain_entities "delivery/Services/Product/Domian/Entities"
)

type OrderAggrigate struct {
	Order order_entities.OrderEntity
	Items []product_domain_entities.ProductEntity
	Price float32
}

func NewOrderAggrigate(order order_entities.OrderEntity, items []product_domain_entities.ProductEntity) *OrderAggrigate {

	orderAgg := &OrderAggrigate{}
	orderAgg.Order = order
	orderAgg.Items = items
	orderAgg.Price = orderAgg.CalculateCoast()

	return orderAgg
}

func (obj *OrderAggrigate) CalculateCoast() float32 {
	coast := float32(0)
	for _, product := range obj.Items {
		coast += product.GetPrice()
	}
	return coast
}
