package authu_services

import (
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"fmt"
	"math/rand"
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
	user, err := obj.repo.Login(dto.Login, shared_utils.Md5Hash(dto.Password))
	if err != nil {
		return err.Error(), err
	}

	token, err := shared_utils.JwtToken(user.Email, user.Id)
	if err != nil {
		return "", fmt.Errorf("could not generate token " + err.Error())
	}

	obj.repo.ClearToken(user.Id)
	auth := obj.repo.CreateAuth(token, user)
	obj.repo.LockUser(auth.Email, "1")
	obj.repo.CleanPins(auth.Email)
	ok, err := obj.repo.TwoFactoryCreate(&auth_infrastructure_models.TwoFactoryPin{
		Pin:   rand.Intn(99999999),
		Email: dto.Login,
	})

	if !ok || err != nil {
		return "", fmt.Errorf("please try again later")
	}
	return token, nil
}

func (obj *AuthService) Register(dto auth_domain_dtos.RegisterDTO) (bool, error) {

	dt := time.Now()
	formattedTime := dt.Format("2006-01-02 15:04:05")
	user := &shared_models.User{
		UserName:  dto.UserName,
		FullName:  dto.FullName,
		Email:     dto.Email,
		Address:   dto.Address,
		Password:  shared_utils.Md5Hash(dto.Password),
		UserType:  "customer",
		UserLevel: dto.UserLevel,
		IsOnline:  0,
		IsLocked:  1,
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
	obj.repo.LockUser(dto.Email, "0")
	return true, nil
}

func (obj *AuthService) Logout(dto auth_domain_dtos.LogoutDTO) {
	obj.repo.Logout(dto.Token)
}
