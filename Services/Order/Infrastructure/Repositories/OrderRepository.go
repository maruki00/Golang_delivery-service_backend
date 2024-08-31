package user_repositories

import (
	orderaggrigate "delivery/Services/Order/Domain/Aggrigates"
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

func (obj *OrderRepository) Make(order *orderaggrigate.OrderAggrigate) (*orderaggrigate.OrderAggrigate, error) {

	

	obj.db.Model().Ins



	res := obj.db.Model(obj.model).Create(&order.OrderEntity)

	if res.Error != nil {
		return nil, errors.New("could not create the order")
	}

	order_product := &OrderProduct

	for _, product := range order.Items {
		ids_products = append(ids_products, product.GetId())
	}

	return nil, nil
}

func (obj *OrderRepository) Delete() error {

	return nil
}

func (obj *OrderRepository) Cancel() error {

	return nil
}

func (obj *OrderRepository) ConfirmePickUp() error {

	return nil
}

func (obj *OrderRepository) Tracking() error {

	return nil
}

func (obj *OrderRepository) Confirm() error {

	return nil
}
