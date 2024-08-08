package authports

import "delivery/Services/Auth/Domain/DTOs"

type RegisterInputPort interface {
	Register(dto DTOs.UserDTO) (string, error)
}
