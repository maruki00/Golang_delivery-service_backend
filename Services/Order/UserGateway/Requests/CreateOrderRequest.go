package order_usergateway_requests

type CreateOrderRequest struct {
	CostomerId int
	Products   [][2]int
}
