package order_infrastructure_repositories

import (
	"context"
	order_domain_aggrigate "delivery/internal/order/Domain/Aggrigates"
	order_domain_entities "delivery/internal/order/Domain/Entities"
	order_infrastructues_models "delivery/internal/order/Infrastructure/Models"
	"errors"
	"fmt"

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

func (obj *OrderRepository) Make(ctx context.Context, order *order_domain_aggrigate.OrderAggrigate) (*order_domain_aggrigate.OrderAggrigate, error) {

	obj.db.Begin()
	res := obj.db.WithContext(ctx).Model(obj.model).Create(order.Order)
	fmt.Println(order.Order)
	if res.Error != nil {
		obj.db.Rollback()
		return nil, errors.New("could not create the order")
	}

	d := obj.db.Model(&order_infrastructues_models.OrderProduct{})
	for _, product := range order.Items {
		res := d.Create(&product)
		if res.Error != nil {
			obj.db.Rollback()
			return nil, errors.New("could not create order prodcut")
		}
	}

	obj.db.Commit()
	return order, nil
}

func (obj *OrderRepository) Cancel(ctx context.Context, id int) error {

	res := obj.db.WithContext(ctx).Where("id = ?", id).Delete(obj.model).Update("is_cancelled", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}

func (obj *OrderRepository) Confirm(ctx context.Context, id int) error {

	res := obj.db.WithContext(ctx).Where("id = ?", id).Update("is_confirmed", 1)
	if res.Error != nil {
		return errors.New("could not delete the order")
	}
	return nil
}

func (obj *OrderRepository) GetStatus(ctx context.Context, id int) (int, error) {
	var order order_infrastructues_models.Order

	res := obj.db.WithContext(ctx).Model(order).Where("id = ?", id).Find(&order)
	if res.Error != nil {
		return -1, res.Error
	}
	return order.Status, nil
}

func (obj *OrderRepository) GetCustomerOrders(ctx context.Context, id int) ([]order_domain_entities.OrderEntity, error) {
	var orders []order_domain_entities.OrderEntity

	res := obj.db.WithContext(ctx).Model(obj.model).Where("customer_id = ?", id).Find(&orders)

	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil

}

func (obj *OrderRepository) GetByFingerPrint(ctx context.Context, fingerprint string) ([]order_domain_entities.OrderEntity, error) {
	orders := []order_domain_entities.OrderEntity{}
	res := obj.db.WithContext(ctx).Model(obj.model).Where("fingerprint = ?", fingerprint).Find(orders)
	if res.Error != nil {
		return nil, res.Error
	}

	return orders, nil
}
