package product_usergetway_requests

type GetProductRequest struct {
	Id int `validate: "required" josn: "id"`
}
