package product_domain_dtos

type UpdateProductDTO struct {
	Id    int
	Label string
	Type  string
	Price float32
}
