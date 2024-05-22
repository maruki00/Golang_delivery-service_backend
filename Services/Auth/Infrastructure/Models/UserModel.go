package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int
	UserName string
	FullName string
	Email    string
}
