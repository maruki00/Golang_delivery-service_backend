package authports

import "delivery/Services/Auth/Domain/DTOs"

type LoginInputPort interface {
	Login(dto DTOs.LoginDTO) (string, error)
}
