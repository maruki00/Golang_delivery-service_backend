package order_infrastructues_models

import (
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	Id               int       `json:"id"`
	OrderFingerprint string    `json:"order_fingerprint"`
	CostumerId       int       `json:"costumer_id"`
	Cost             float32   `json: "cost"`
	Status           int       `json: "status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (obj *OrderModel) GetId() int {
	return obj.Id
}
func (obj *OrderModel) GetOrderFingerprint() string {
	return obj.OrderFingerprint
}
func (obj *OrderModel) GetCostumerId() int {
	return obj.CostumerId
}
func (obj *OrderModel) GetCost() float32 {
	return obj.Cost
}
func (obj *OrderModel) GetStatus() int {
	return obj.Status
}
