package authu_services

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
)

type AuthService struct {
	repo *auth_infrastructure_repository.AuthRepository
}

func NewAuthService(repo *auth_infrastructure_repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (obj *AuthService) Login(dto auth_domain_dtos.LoginDTO) (string, error) {
	accessToken, err := obj.repo.Login(dto.GetLogin(), dto.GetPassword())
	if err != nil {
		return err.Error(), err
	}
	return accessToken, nil
}
