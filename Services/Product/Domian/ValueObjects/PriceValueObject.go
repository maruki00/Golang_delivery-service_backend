package product_domain_valueobjects

import (
	"fmt"
)

type PriceVauleObject struct {
	price float32
}

func (p *PriceVauleObject) Set(price float32) {
	p.price = price
}

func (p *PriceVauleObject) Get(price float32) float32 {
	return p.price
}

func (p *PriceVauleObject) ToStriong(price float32) string {
	return fmt.Sprintf("%f", p.price)
}
