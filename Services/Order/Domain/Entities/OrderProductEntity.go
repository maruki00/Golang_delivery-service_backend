package order_domain_entities

type OrderProductEntity interface {
	GetId() int
	GetOrderId() int
	GetProductId() int
	GetQty() int
}
