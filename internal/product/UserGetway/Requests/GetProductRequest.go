package product_usergateway_requests

type GetProductRequest struct {
	Id int `validate: "required" josn: "id"`
}
