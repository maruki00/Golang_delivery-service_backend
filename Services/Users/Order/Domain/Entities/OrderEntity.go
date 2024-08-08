package order_entities

import (
	"gorm.io/gorm"
)

type OrderEntity interface {
}

type OrderModel struct {
	gorm.Model
}
