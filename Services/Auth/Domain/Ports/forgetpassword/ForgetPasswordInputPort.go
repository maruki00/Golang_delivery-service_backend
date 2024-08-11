package auth_domain_ports

import auth_domain_dto "delivery/Services/Auth/Domain/DTOs"

type ForgetPasswordInputPort interface {
	Handel(dto auth_domain_dto.FrogetPasswordDTO) (string, error)
}
