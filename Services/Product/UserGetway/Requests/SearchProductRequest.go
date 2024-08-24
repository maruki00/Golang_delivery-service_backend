package product_usergetway_requests

type SearchProductRequest struct {
	Query string `validate : "required" json: "query"`
}
