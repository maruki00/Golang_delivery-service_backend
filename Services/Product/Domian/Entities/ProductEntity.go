package product_domain_entities

type ProductEntity interface {
	GetLabel() string
	GetId() int
	GetPrice() float32
	GetType() string
}
