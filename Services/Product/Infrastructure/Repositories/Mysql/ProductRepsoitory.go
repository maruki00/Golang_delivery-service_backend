package mysql_product_repository

import (
	"crypto/md5"

	product_domain_entities "delivery/Services/Product/Domian/Entities"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	shareddb "delivery/Services/Shared/Infrastructure/DB"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(dbObj *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: dbObj,
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

func (obj *ProductRepository) Delete(login string, password string) error {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &DTOs.LoginDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
	if dto.Failed() {
		return errors.New("invalid credential")
	}
	return nil
}

func (obj *ProductRepository) GetById(login string, password string) error {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &DTOs.LoginDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
	if dto.Failed() {
		return errors.New("invalid credential")
	}
	return nil
}

func (obj *ProductRepository) Search(login string, password string) error {
	db := shareddb.NewDB()
	defer db.Close()
	dto := &DTOs.LoginDTO{}

	hash := md5.Sum([]byte(password))

	h := fmt.Sprintf("%x", hash)

	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
	defer statement.Close()
	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
	if dto.Failed() {
		return errors.New("invalid credential")
	}
	return nil
}
