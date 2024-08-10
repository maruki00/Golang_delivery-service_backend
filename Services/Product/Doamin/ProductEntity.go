package product_domain

type ProductEntity interface {
	GetId() int
	GetLabale() string
	GetPrice() float32
	SetId(id int)
	SetLabale(label string)
	SetPrice(GetPrice float32)
}
