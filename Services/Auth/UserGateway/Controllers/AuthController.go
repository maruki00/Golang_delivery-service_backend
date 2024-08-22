package auth_usergetway_controllers

import (
	authu_services "delivery/Services/Auth/Application/Services"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	auth_requests "delivery/Services/Auth/UserGateway/Requests"
	shared_configs "delivery/Services/Shared/Application/Configs"
	shared_utils "delivery/Services/Shared/Application/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	Validate *validator.Validate
	service  *authu_services.AuthService
}

func NewAuthController() *AuthController {
	config, _ := shared_configs.GetConfig()
	return &AuthController{
		Validate: validator.New(),
		service:  authu_services.NewAuthService(auth_infrastructure_repository.NewAuthRepository(config)),
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &auth_requests.LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	accessToken, _ := obj.service.Login(auth_domain_dtos.LoginDTO{
		Login:    request.Login,
		Password: request.Password,
	})

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{
		"message": "Success",
		"token":   accessToken,
		"error":   nil,
		"data":    nil,
	})
}

func (obj AuthController) Register(ctx *gin.Context) {
	request := &auth_requests.RegisterRequest{}
	if err := ctx.BindJSON(request); err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	_, err := obj.service.Register(auth_domain_dtos.RegisterDTO{
		FullName: request.FullName,
		UserName: request.UserName,
		Email:    request.Email,
		Address:  request.UserName,
		Password: request.Password,
	})

	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", nil)
}

func (obj AuthController) TwoFactoryConfirm(ctx *gin.Context) {
	request := &auth_requests.TwoFactoryConfirmRequest{}
	if err := ctx.BindJSON(request); err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	// _, err := obj.service.Register(auth_domain_dtos.RegisterDTO{
	// 	FullName: request.FullName,
	// 	UserName: request.UserName,
	// 	Email:    request.Email,
	// 	Address:  request.UserName,
	// 	Password: request.Password,
	// })

	// if err != nil {
	// 	shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
	// 	return
	// }

	shared_utils.Success(ctx, http.StatusOK, "Success", nil)
}
