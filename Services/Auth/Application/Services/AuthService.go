package authu_services

import (
	domain_auth_contracts "delivery/Services/Auth/Domain/Contracts"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_domain_ports "delivery/Services/Auth/Domain/Ports"
	auth_infrastructure_models "delivery/Services/Auth/Infrastructure/Models"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repo    *auth_infrastructure_repository.AuthRepository
	outport auth_domain_ports.AuthOutputPort
}

func NewAuthService(repo *auth_infrastructure_repository.AuthRepository, outport auth_domain_ports.AuthOutputPort) *AuthService {
	return &AuthService{
		repo:    repo,
		outport: outport,
	}
}

func (obj *AuthService) Login(dto auth_domain_dtos.LoginDTO) domain_auth_contracts.ViewModel {

	user, err := obj.repo.Login(dto.Login, shared_utils.Md5Hash(dto.Password))
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Data:    nil,
		})
	}

	token, err := shared_utils.JwtToken(user.Email, user.Id)
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Data:    nil,
		})
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
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Data:    nil,
		})
	}
	return obj.outport.Success(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Data:    gin.H{"result": token},
	})
}

func (obj *AuthService) Register(dto auth_domain_dtos.RegisterDTO) domain_auth_contracts.ViewModel {

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
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Data:    nil,
		})
	}
	return obj.outport.Error(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Data:    nil,
	})
}

func (obj *AuthService) TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) domain_auth_contracts.ViewModel {

	_, err := obj.repo.TwoFactoryConfirm(dto.Email, dto.Pin)
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Data:    nil,
		})
	}
	obj.repo.LockUser(dto.Email, "0")
	return obj.outport.Success(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Data:    nil,
	})
}

func (obj *AuthService) Logout(dto auth_domain_dtos.LogoutDTO) {
	obj.repo.Logout(dto.Token)
}
