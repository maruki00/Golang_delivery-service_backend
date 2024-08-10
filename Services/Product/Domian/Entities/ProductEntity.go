package product_domain_entities

type ProductEntity interface {
	GetId() int
	GetLabale() string
	GetPrice() float32
	SetId(id int)
	SetLabale(label string)
	SetPrice(GetPrice float32)
}
