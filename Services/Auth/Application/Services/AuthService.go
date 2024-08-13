package authu_services

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_domain_login_ports "delivery/Services/Auth/Domain/Ports/login"
)

type AuthService struct {
	repo    domain_auth_contracts.IAuthRepository
	outPort auth_domain_login_ports.LoginOutputPort
}

func NewAuthService(repo domain_auth_contracts.IAuthRepository, outputPort auth_domain_login_ports.LoginOutputPort) *AuthService {
	return &AuthService{
		repo:    repo,
		outPort: outputPort,
	}
}

func (obj AuthService) Login(dto auth_domain_dtos.LoginDTO) (string, error) {

	_, err := obj.repo.Login(dto.GetLogin(), dto.GetPassword())
	if err != nil {
		return obj.outPort.Success(), nil
	}
	return "success", nil
}
