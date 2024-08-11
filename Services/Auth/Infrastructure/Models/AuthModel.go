package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	username string                          `json:user_name`
	fullname string                          `json:full_name`
	email    ValueObject.EmailValueObject    `json:email`
	password ValueObject.PasswordValueObject `json:password`
	address  ValueObject.AddressValueObject  `json: address`
	userType string                          `json: user_type`
}

func (obj *UserModel) GetUserName() string {
	return obj.username
}
func (obj *UserModel) GetFullName() string {
	return obj.fullname
}
func (obj *UserModel) GetEmail() ValueObject.EmailValueObject {
	return obj.email
}
func (obj *UserModel) GetPassword() ValueObject.PasswordValueObject {
	return obj.password
}
func (obj *UserModel) GetAddress() ValueObject.AddressValueObject {
	return obj.address
}
func (obj *UserModel) SetUserName(username string) {
	obj.username = username
}
func (obj *UserModel) SetFullName(fullName string) {
	obj.fullname = fullName
}
func (obj *UserModel) SetEmail(email ValueObject.EmailValueObject) {
	obj.email = email
}
func (obj *UserModel) SetPassword(password ValueObject.PasswordValueObject) {
	obj.password = password
}
func (obj *UserModel) SetAddress(address ValueObject.AddressValueObject) {
	obj.address = address
}