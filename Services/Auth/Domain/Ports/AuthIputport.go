package auth_domain_ports

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
)

type AuthInputPort interface {
	Login(dto auth_domain_dtos.LoginDTO) shared_domain_contracts.ViewModel
	Register(dto auth_domain_dtos.RegisterDTO) shared_domain_contracts.ViewModel
	TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) shared_domain_contracts.ViewModel
	Logout(dto auth_domain_dtos.LogoutDTO)
}
