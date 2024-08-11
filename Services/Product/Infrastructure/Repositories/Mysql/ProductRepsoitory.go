// package mysql_product_repository

// import (
// 	"crypto/md5"

// 	shareddb "delivery/Services/Shared/Infrastructure/DB"
// 	"errors"
// 	"fmt"
// )

// type ProductRepository struct {
// }

// func NewProductRepository() *ProductRepository {
// 	return &ProductRepository{}
// }

// func (obj *ProductRepository) Insert(label string, price PriceVauleObject) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }

// func (obj *ProductRepository) Insert(login string, password string) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }

// func (obj *ProductRepository) Update(login string, password string) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }

// func (obj *ProductRepository) Delete(login string, password string) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }

// func (obj *ProductRepository) GetById(login string, password string) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }

// func (obj *ProductRepository) Search(login string, password string) error {
// 	db := shareddb.NewDB()
// 	defer db.Close()
// 	dto := &DTOs.LoginDTO{}

// 	hash := md5.Sum([]byte(password))

// 	h := fmt.Sprintf("%x", hash)

// 	statement, _ := db.Prepare("SELECT email, password FROM users WHERE email = ? and password = ? limit 1")
// 	defer statement.Close()
// 	statement.QueryRow(login, h).Scan(&dto.Login, &dto.Password)
// 	if dto.Failed() {
// 		return errors.New("invalid credential")
// 	}
// 	return nil
// }
