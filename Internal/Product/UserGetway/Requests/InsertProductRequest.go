package product_usergateway_requests

type InsertProductRequest struct {
	Label string  `validate:"required" json:"label"`
	Price float32 `validate:"required" json:"price"`
	Type  string  `validate:"required" json:"type"`
}
