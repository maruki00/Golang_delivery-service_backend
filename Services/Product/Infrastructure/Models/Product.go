package product_infrastructure_models

type Product struct {
	Id    int
	Label string
	Price float32
}

func (p *Product) GetId() int {
	return p.Id
}
func (p *Product) GetLabale() string {
	return p.Label
}
func (p *Product) GetPrice() float32 {
	return p.Price
}

func (p *Product) SetId(id int) {
	p.Id = id
}
func (p *Product) SetLabel(label string) {
	p.Label = label
}
func (p *Product) SetPrice(price float32) {
	p.Price = price
}
