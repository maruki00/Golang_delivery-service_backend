package mysql_product_repository

import (
	"crypto/md5"

	product_domain_entities "delivery/Services/Product/Domian/Entities"
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

func (obj *ProductRepository) Insert(product product_domain_entities.ProductEntity) error {
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

func (obj *ProductRepository) Update(login string, password string) error {
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
