package product_usergateway_requests

type SearchProductRequest struct {
	Query string `validate : "required" json: "query"`
}
