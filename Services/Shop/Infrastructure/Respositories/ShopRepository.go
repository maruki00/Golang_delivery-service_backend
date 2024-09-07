package shop_infrastructure_respositories

import "gorm.io/gorm"

type ShopRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewShopRepository(db *gorm.DB, model interface{}) *ShopRepository {
	return &ShopRepository{
		db:    db,
		model: model,
	}
}

func (obj *ShopRepository) SetStatus(id int, status int) error {

	res := obj.db.Model(obj.model).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	return nil

}
