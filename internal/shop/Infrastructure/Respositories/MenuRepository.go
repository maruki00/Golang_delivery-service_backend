package shop_infrastructure_respositories

import (
	shop_domain_entities "delivery/internal/shop/Domain/Entities"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db    *gorm.Model
	model interface{}
}

func NewMenuRepository(db *gorm.Model, model interface{}) *MenuRepository {
	return &MenuRepository{
		db:    db,
		model: model,
	}
}

func (obj *MenuRepository) Make(menn shop_domain_entities.MenuEntity) (shop_domain_entities.MenuEntity, error) {

	// res := obj.

	return nil, nil
}
