package authu_services

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
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
	accessToken, err := obj.repo.Login(dto.Login, dto.Password)
	if err != nil {
		return err.Error(), err
	}
	return accessToken, nil
}

func (obj *AuthService) Register(dto auth_domain_dtos.RegisterDTO) (bool, error) {

	user := *&shared_models.User{
		Username:  dto.UserName,
		Fullname:  dto.FullName,
		Email:     dto.Email,
		Address:   dto.Address,
		Password:  dto.Password,
		UserType:  dto.UserType,
		UserLevel: dto.UserLevel,
		IsOnline:  "0",
		IsLocked:  "0",
	}
	err := obj.repo.Register(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
