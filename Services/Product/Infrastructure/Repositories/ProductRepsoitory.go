package product_Infrastructure_repository

import (
	product_domain_entities "delivery/Services/Product/Domian/Entities"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewProductRepository(dbObj *gorm.DB, model interface{}) *ProductRepository {
	return &ProductRepository{
		db:    dbObj,
		model: model,
	}
}

func (obj *ProductRepository) Insert(product product_domain_entities.ProductEntity) (product_domain_entities.ProductEntity, error) {
	res := obj.db.Model(&product_infrastructure_models.Product{}).Create(product)
	fmt.Println("product : ", product)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("record could not be saved")
	}
	return product, nil
}

func (obj *ProductRepository) GetById(id int) (product_domain_entities.ProductEntity, error) {

	product := &product_infrastructure_models.Product{}
	obj.db.Model(&product_infrastructure_models.Product{}).Where("id = ?", id).Find(&product)
	if product == nil {
		return nil, fmt.Errorf("record could not be found")
	}
	if product.Id == 0 {
		return nil, nil
	}

	return product, nil
}

func (obj *ProductRepository) Search(seasrch string) ([]product_infrastructure_models.Product, error) {
	items := []product_infrastructure_models.Product{}
	res := obj.db.Model(obj.model).Where("( label like  ?  or price like  ?  or type like  ? )", fmt.Sprintf("%%%s", seasrch), fmt.Sprintf("%%%s%%", seasrch), fmt.Sprintf("%%%s%%", seasrch)).Limit(3).Offset(0).Find(&items)
	if res.Error != nil {
		return []product_infrastructure_models.Product{}, nil
	}
	return items, nil
}

func (obj *ProductRepository) Update(id int, data map[string]interface{}) (product_domain_entities.ProductEntity, error) {

	fmt.Println("result : ", data)
	res := obj.db.Model(&obj.model).Where("id", id).Updates(data)
	if res.Error != nil {
		return nil, fmt.Errorf("something happen, %v", res.Error)
	}

	product, err := obj.GetById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (obj *ProductRepository) Delete(id int) (product_domain_entities.ProductEntity, error) {

	product, err := obj.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("product not found")
	}
	res := obj.db.Model(obj.model).Where("id = ? ", id).Delete(&product)
	if res.Error != nil {
		return nil, fmt.Errorf("something happen %v", res.Error)
	}
	if res.RowsAffected == 0 {
		return nil, fmt.Errorf("could not delete the record")
	}
	return product, nil
}
