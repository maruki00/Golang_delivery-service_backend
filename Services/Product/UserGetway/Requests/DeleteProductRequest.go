package product_usergetway_requests

type DeleteProductRequest struct {
	Id int `validate: "required" josn: "id"`
}
