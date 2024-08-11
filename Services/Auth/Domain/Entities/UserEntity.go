package Entities

type AuthEntity interface {
	GetEmail() string
	GetPassowrd() string

	SetEmail(email string)
	SetPassword(password string)
}
