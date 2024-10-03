package product_usergateway_requests

type MultipleProductstRequest struct {
	Ids string `validate: "required" josn: "ids"`
}
