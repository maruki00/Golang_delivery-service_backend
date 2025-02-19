package controllers

import (
	"delivery/internal/auth/userGateWay/requests"
	shared_utils "delivery/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthController struct {
	Validate *validator.Validate

	inputPort sahred_domain.AuthInputPort
}

func NewAuthController(db *gorm.DB) *AuthController {
	repo := auth_infra_repository.NewAuthRepository(db)
	presenter := &auth_usergateway_presenters.JsonAuthPresenter{}

	return &AuthController{
		Validate:  validator.New(),
		inputPort: auth_services.NewAuthService(repo, presenter),
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &requests.LoginRequest{}
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

	res := obj.inputPort.Register(auth_domain_dtos.RegisterDTO{
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

	res := obj.inputPort.TwoFactoryConfirm(auth_domain_dtos.TwoFactoryConfirmDTO{
		Email: request.Email,
		Pin:   request.Pin,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())

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
	obj.inputPort.Logout(auth_domain_dtos.LogoutDTO{
		Token: request.Token,
	})
	shared_utils.Success(ctx, http.StatusNoContent, "Success", nil)
}
