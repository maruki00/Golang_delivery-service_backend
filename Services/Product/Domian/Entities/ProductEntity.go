package product_domain_entities

type ProductEntity interface {
	GetLabel() string
	GetPrice() float32
	GetType() string
}
