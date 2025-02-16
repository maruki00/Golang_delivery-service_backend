package product_usergateway_requests

type DeleteProductRequest struct {
	Id int `validate: "required" josn: "id"`
}
