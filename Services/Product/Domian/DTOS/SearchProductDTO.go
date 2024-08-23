package product_domain_dtos

type SearchProductDTO struct {
	id    int64
	query string
}

func (dto SearchProductDTO) GetId() int64 {
	return dto.id
}

func (dto SearchProductDTO) GetQuery() string {
	return dto.query
}
