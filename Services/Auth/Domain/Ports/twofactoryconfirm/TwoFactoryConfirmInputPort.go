package authports

import auth_domain_dto "delivery/Services/Auth/Domain/DTOs"

type TwoFactoryConfirmInputPort interface {
	Handel(dto auth_domain_dto.TwoFactoryConfirmDTO) (string, error)
}
