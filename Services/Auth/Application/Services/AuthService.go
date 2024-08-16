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

func Login(dto auth_domain_dtos.LoginDTO) (string, error) {
	_, _ = repo.Login(dto.GetLogin(), dto.GetPassword())
	// if err != nil {
	// 	return "invalid credentials", nil
	// }
	return "success", nil
}
