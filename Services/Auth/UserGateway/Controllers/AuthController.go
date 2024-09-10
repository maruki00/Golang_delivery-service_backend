package auth_usergetway_controllers

import (
	authu_services "delivery/Services/Auth/Application/Services"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	auth_domain_ports "delivery/Services/Auth/Domain/Ports"
	auth_infrastructure_repository "delivery/Services/Auth/Infrastructure/Repositories"
	auth_requests "delivery/Services/Auth/UserGateway/Requests"
	shared_utils "delivery/Services/Shared/Application/Utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthController struct {
	Validate  *validator.Validate
	service   *authu_services.AuthService
	inputPort auth_domain_ports.AuthInputPort
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		Validate:  validator.New(),
		service:   authu_services.NewAuthService(auth_infrastructure_repository.NewAuthRepository(db), nil),
		inputPort: nil,
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

	result := obj.inputPort.Login(auth_domain_dtos.LoginDTO{
		Login:    request.Login,
		Password: request.Password,
	})

	ctx.JSON(result.GetResponse().Status, result.GetResponse())
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

	res := obj.service.Register(auth_domain_dtos.RegisterDTO{
		FullName:  request.FullName,
		UserName:  request.UserName,
		Email:     request.Email,
		Address:   request.Address,
		Password:  request.Password,
		UserLevel: 0,
		UserType:  "costumer",
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())

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

	res := obj.service.TwoFactoryConfirm(auth_domain_dtos.TwoFactoryConfirmDTO{
		Email: request.Email,
		Pin:   request.Pin,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())

	// if err != nil || !res {
	// 	shared_utils.Error(ctx, http.StatusInternalServerError, "Error", "could not confirm the account")
	// 	return
	// }

	// shared_utils.Success(ctx, http.StatusOK, "Success", nil)
}

func (obj *AuthController) Logout(ctx *gin.Context) {
	request := &auth_requests.LogoutRequest{}
	if err := ctx.BindJSON(request); err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		shared_utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}
	obj.service.Logout(auth_domain_dtos.LogoutDTO{
		Token: request.Token,
	})
	shared_utils.Success(ctx, http.StatusNoContent, "Success", nil)
}
