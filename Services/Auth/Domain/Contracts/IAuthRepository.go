package domain_auth_contracts

import auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"

type IAuthRepository interface {
	Login(login, password string) (*auth_domain_dtos.AuthDTO, error)
}
