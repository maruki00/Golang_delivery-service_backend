package order_entities

import (
	"delivery/Auth/Domain/ValueObject"

	"gorm.io/gorm"
)

type OrderEntity interface {
	GetUserName() string
	GetFullName() string
	GetEmail() ValueObject.EmailValueObject
	GetPassword() ValueObject.PasswordValueObject
	GetAddress() ValueObject.AddressValueObject

	SetUserName(username string)
	SetFullName(fullName string)
	SetEmail(email ValueObject.EmailValueObject)
	SetPassword(password ValueObject.PasswordValueObject)
	SetAddress(address ValueObject.AddressValueObject)
	SetType(Type string)
}

type OrderModel struct {
	gorm.Model
	username string                          `json:user_name`
	fullname string                          `json:full_name`
	email    ValueObject.EmailValueObject    `json:email`
	password ValueObject.PasswordValueObject `json:password`
	address  ValueObject.AddressValueObject  `json: address`
	userType string                          `json: user_type`
}

func (obj *OrderModel) GetUserName() string {
	return obj.username
}
func (obj *OrderModel) GetFullName() string {
	return obj.fullname
}
func (obj *OrderModel) GetEmail() ValueObject.EmailValueObject {
	return obj.email
}
func (obj *OrderModel) GetPassword() ValueObject.PasswordValueObject {
	return obj.password
}
func (obj *OrderModel) GetAddress() ValueObject.AddressValueObject {
	return obj.address
}
func (obj *OrderModel) SetUserName(username string) {
	obj.username = username
}
func (obj *OrderModel) SetFullName(fullName string) {
	obj.fullname = fullName
}
func (obj *OrderModel) SetEmail(email ValueObject.EmailValueObject) {
	obj.email = email
}
func (obj *OrderModel) SetPassword(password ValueObject.PasswordValueObject) {
	obj.password = password
}
func (obj *OrderModel) SetAddress(address ValueObject.AddressValueObject) {
	obj.address = address
}
