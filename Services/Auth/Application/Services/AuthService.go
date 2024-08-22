package authu_services

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"fmt"
	"time"
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

	twoFactory, err := obj.repo.TwoFactoryCreate(&auth_infrastructure_models.TwoFactoryPin{
		Pin:   12334,
		Email: dto.Login
	})
	return accessToken, nil
}

func (obj *AuthService) Register(dto auth_domain_dtos.RegisterDTO) (bool, error) {

	dt := time.Now()
	formattedTime := dt.Format("2006-01-02 15:04:05")
	fmt.Println("now : ", formattedTime, dt)
	user := &shared_models.User{
		UserName:  dto.UserName,
		FullName:  dto.FullName,
		Email:     dto.Email,
		Address:   shared_utils.Md5Hash(dto.Address),
		Password:  dto.Password,
		UserType:  "customer",
		UserLevel: dto.UserLevel,
		IsOnline:  0,
		IsLocked:  0,
		LastSeen:  formattedTime,
		CreatedAt: formattedTime,
		UpdatedAt: formattedTime,
	}
	_, err := obj.repo.Register(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (obj *AuthService) TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) (bool, error) {

	_, err := obj.repo.TwoFactoryConfirm(dto.Email, dto.Pin)
	if err != nil {
		return false, err
	}
	return true, nil
}
