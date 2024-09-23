package product_usergetway_requests

type MultipleProductstRequest struct {
	Ids string `validate: "required" josn: "ids"`
}
