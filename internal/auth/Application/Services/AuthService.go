package auth_services

import (
	auth_domain_contracts "delivery/internal/auth/Domain/Contracts"
	auth_domain_dtos "delivery/internal/auth/Domain/DTOs"
	auth_domain_ports "delivery/internal/auth/Domain/Ports"
	auth_infrastructure_models "delivery/internal/auth/Infrastructure/Models"
	shared_utils "delivery/internal/shared/Application/Utils"
	shared_domain_contracts "delivery/internal/shared/Domain/Contracts"
	shared_models "delivery/internal/shared/Infrastructure/Models"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repo    auth_domain_contracts.IAuthRepository
	outport auth_domain_ports.AuthOutputPort
}

func NewAuthService(repo auth_domain_contracts.IAuthRepository, outport auth_domain_ports.AuthOutputPort) *AuthService {
	return &AuthService{
		repo:    repo,
		outport: outport,
	}
}

func (obj *AuthService) Login(dto auth_domain_dtos.LoginDTO) shared_domain_contracts.ViewModel {

	user, err := obj.repo.Login(dto.Login, shared_utils.Md5Hash(dto.Password))
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	token, err := shared_utils.JwtToken(user.Email, user.Id)
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
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
			Result:  nil,
		})
	}
	return obj.outport.Success(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"result": token},
	})
}

func (obj *AuthService) Register(dto auth_domain_dtos.RegisterDTO) shared_domain_contracts.ViewModel {

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
			Result:  nil,
		})
	}
	return obj.outport.Error(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Result:  nil,
	})
}

func (obj *AuthService) TwoFactoryConfirm(dto auth_domain_dtos.TwoFactoryConfirmDTO) shared_domain_contracts.ViewModel {

	_, err := obj.repo.TwoFactoryConfirm(dto.Email, dto.Pin)
	if err != nil {
		return obj.outport.Error(shared_models.ResponseModel{
			Status:  400,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}
	obj.repo.LockUser(dto.Email, "0")
	return obj.outport.Success(shared_models.ResponseModel{
		Status:  200,
		Message: "Success",
		Error:   nil,
		Result:  nil,
	})
}

func (obj *AuthService) Logout(dto auth_domain_dtos.LogoutDTO) {
	obj.repo.Logout(dto.Token)
}
