package mysql_product_repository

import (
	product_domain_entities "delivery/Services/Product/Domian/Entities"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
	model interface{}
}

func NewProductRepository(dbObj *gorm.DB, model interface{}) *ProductRepository {
	return &ProductRepository{
		db: dbObj,
		model: model,
	}
}

func (obj *ProductRepository) Insert(product product_domain_entities.ProductEntity) (bool, error) {

	res := obj.db.Model(&product_infrastructure_models.Product{}).Create(product)
	if res.Error != nil {
		return false, res.Error
	}
	if res.RowsAffected == 0 {
		return false, errors.New("record could not be saved")
	}
	return true, nil
}

func (obj *ProductRepository) GetById(id int) (product_domain_entities.ProductEntity, error) {
	var product product_domain_entities.ProductEntity
	obj.db.Model(&product_infrastructure_models.Product{}).Where("id = ?", id).Find(&product)
	if product == nil {
		return nil, fmt.Errorf("record could not be found")
	}

	return product, nil
}

func (obj *ProductRepository) Search(seasrch string) ([]*product_domain_entities.ProductEntity, error) {
	var items []*product_domain_entities.ProductEntity
	res := obj.db.Model(&product_infrastructure_models.Product{}).Where("label = ? or price = ? or type = ? ", seasrch, seasrch, seasrch).Find(&items)
	if res.Error != nil {
		return []*product_domain_entities.ProductEntity{}, nil
	}
	return items, nil
}

func (obj *ProductRepository) Update(id int, data map[string]interface{}) (bool, error) {

	res := obj.db.Model(&product_infrastructure_models.Product{}).Where("id = ?").Updates(data)
	if res.Error != nil {
		return false, fmt.Errorf("Something happen, %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return false, fmt.Errorf("could not update the record")
	}
	return true, nil
}

func (obj *ProductRepository) Delete(id int) (bool, error) {

	res := obj.db.Model(&)
	return true, nil
}
