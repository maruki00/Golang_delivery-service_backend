package orderaggrigate

import (
	order_entities "delivery/Services/Order/Domain/Entities"
	product_domain_entities "delivery/Services/Product/Domian/Entities"
)

type OrderAggrigate struct {
	OrderEntity order_entities.OrderEntity
	Items       []product_domain_entities.ProductEntity
	Price       float32
}

func NewOrderAggrigate(order order_entities.OrderEntity, items []product_domain_entities.ProductEntity) *OrderAggrigate {
	price := float32(0)
	for _, product := range items {
		price += product.GetPrice()
	}

	return &OrderAggrigate{
		OrderEntity: order,
		Items:       items,
		Price:       price,
	}

}
