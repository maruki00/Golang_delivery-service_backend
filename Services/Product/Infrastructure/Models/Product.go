package product_infrastructure_models

type Product struct {
	Label string
	Price float32
	Type  string
}

func (p *Product) GetLabel() string {
	return p.Label
}
func (p *Product) GetPrice() float32 {
	return p.Price
}
func (p *Product) GetType() string {
	return p.Type
}
