package user_repositories

import "gorm.io/gorm"

type OrderRepository struct {
	db *gorm.DB
}

func (obj *OrderRepository) Make(order orderAgriggate) error {

	return nil
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

}

func (obj *OrderRepository) Confirm() error {

}
