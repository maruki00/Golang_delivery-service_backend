package order_infrastructues_models

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	Id        int
	OrderId   int
	ProductId int
}
