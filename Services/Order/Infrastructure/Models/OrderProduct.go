package order_infrastructues_models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	Id        int `json: "id"`
	OrderId   int `json: "order_id"`
	ProductId int `json: "product_id"`
	Qty       int `json: "qty"`
}
