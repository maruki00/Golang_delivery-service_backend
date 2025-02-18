package order_domain_entities

type OrderEntity interface {
	GetId() int
	GetOrderFingerprint() string
	GetCostumerId() int
	GetCost() float32
	GetStatus() int
}
