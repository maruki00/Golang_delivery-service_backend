package models

import "gorm.io/gorm"

type ResetPassword struct {
	gorm.Model
	Id    int
	Email string
	Pin   int
	Token string
	// CreatedAt string
	// UpdatedAt string
}

func (obj *ResetPassword) GetId() int {
	return obj.Id
}
func (obj *ResetPassword) GetEmail() string {
	return obj.Email
}
func (obj *ResetPassword) GetToken() string {
	return obj.Token
}
