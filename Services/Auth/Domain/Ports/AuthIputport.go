package auth_domain_ports

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
)

type AuthInputPort interface {
	Login(dto auth_domain_dtos.LoginDTO) domain_auth_contracts.ViewModel
	Register(dto auth_domain_dtos.RegisterDTO) domain_auth_contracts.ViewModel
	TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) domain_auth_contracts.ViewModel
	Logout(dto auth_domain_dtos.LogoutDTO) domain_auth_contracts.ViewModel
}
