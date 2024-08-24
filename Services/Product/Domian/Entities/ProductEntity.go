package product_domain_entities

type ProductEntity interface {
	GetId() int
	GetLabel() string
	GetPrice() float32
	GetType() string
}
