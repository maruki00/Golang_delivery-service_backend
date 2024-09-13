package order_domain_dtos

type CreateNewOrderDTO struct {
	CostomerId int
	Products   map[int]int
}
