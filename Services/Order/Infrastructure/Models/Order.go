package order_infrastructues_models

import "gorm.io/gorm"

type OrderModel struct {
	gorm.Model
	Id               int     `json:"id"`
	OrderFingerprint string  `json:"order_fingerprint"`
	CostumerId       int     `json:"costumer_id"`
	FromLong         float32 `json:"from_long"`
	FromLat          float32 `json:"from_lat"`
	ToLong           float32 `json:"to_long"`
	ToLat            float32 `json:"to_lat"`
	// CreatedAt        `json:"created_at"`
	// UpdatedAt        timestamp `json:"updated_at"`
}
