package Entities

import (
	"delivery/Auth/Domain/ValueObject"
)

type UserEntity interface {
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
