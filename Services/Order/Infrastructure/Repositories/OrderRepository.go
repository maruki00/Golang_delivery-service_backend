package user_repositories

import (
	order_domain_entities "delivery/Services/Order/Domain/Entities"
	"errors"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewOrderRepository(db *gorm.DB, model interface{}) *OrderRepository {
	return &OrderRepository{
		db:    db,
		model: model,
	}
}

func (obj *OrderRepository) Make(order order_domain_entities.OrderEntity) (order_domain_entities.OrderEntity, error) {

	res := obj.db.Model(obj.model).Create(&order)

	if res.Error != nil {
		return nil, errors.New("could not create the order")
	}

	return order, nil
}

func (obj *OrderRepository) CreateOrderProducts(model interface{}, products []order_domain_entities.OrderProductEntity) ([]order_domain_entities.OrderProductEntity, error) {

	obj.db.Begin()
	d := obj.db.Model(model)
	for _, product := range products {
		res := d.Create(product)
		if res.Error != nil {
			obj.db.Rollback()
			return nil, errors.New("could not add a new record")
		}
	}
	obj.db.Commit()
	return nil, nil
}

func (obj *OrderRepository) Cancel(id int) error {

	res := obj.db.Where("id = ?", id).Delete(obj.model).Update("is_cancelled", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}

func (obj *OrderRepository) Confirm(id int) error {

	res := obj.db.Where("id = ?", id).Update("is_confirmed", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}
