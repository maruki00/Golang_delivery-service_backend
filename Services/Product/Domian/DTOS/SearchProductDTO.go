package product_domain_dtos

import (
	"errors"
	"strconv"
)

type SearchProductDTO struct {
	id    int64
	query string
}

func SearchProductDTOFromArray(array map[string]string) (*SearchProductDTO, error) {
	id, ok := array["id"]
	if !ok {
		return nil, errors.New("key id is missing")
	}
	query, ok := array["query"]
	if !ok {
		return nil, errors.New("key query is missing")
	}
	idVal, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("id should be integer")
	}
	return &SearchProductDTO{
		id:    int64(idVal),
		query: query,
	}, nil
}

func (dto SearchProductDTO) GetId() int64 {
	return dto.id
}

func (dto SearchProductDTO) GetQuery() string {
	return dto.query
}
