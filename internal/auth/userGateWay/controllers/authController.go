package controllers

import (
	"delivery/internal/auth/app/services"
	"delivery/internal/auth/domain/dtos"
	"delivery/internal/auth/domain/ports"
	"delivery/internal/auth/infra/repositories"
	"delivery/internal/auth/userGateWay/adapters/presenters"
	"delivery/internal/auth/userGateWay/requests"
	"delivery/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AuthController struct {
	Validate *validator.Validate

	inputPort ports.AuthInputPort
}

func NewAuthController(db *gorm.DB) *AuthController {
	repo := repositories.NewAuthRepository(db)
	presenter := &presenters.JsonAuthPresenter{}

	return &AuthController{
		Validate:  validator.New(),
		inputPort: services.NewAuthService(repo, presenter),
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &requests.LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	result := obj.inputPort.Login(dtos.LoginDTO{
		Login:    request.Login,
		Password: request.Password,
	})

	ctx.JSON(result.GetResponse().Status, result.GetResponse())
}

func (obj AuthController) Register(ctx *gin.Context) {

	request := &requests.RegisterRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	res := obj.inputPort.Register(dtos.RegisterDTO{
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
	request := &requests.TwoFactoryConfirmRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}

	res := obj.inputPort.TwoFactoryConfirm(dtos.TwoFactoryConfirmDTO{
		Email: request.Email,
		Pin:   request.Pin,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())

}

func (obj *AuthController) Logout(ctx *gin.Context) {
	request := &requests.LogoutRequest{}
	if err := ctx.BindJSON(request); err != nil {
		utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		utils.Error(ctx, http.StatusBadRequest, "Bad Request", errorMessage)
		return
	}
	obj.inputPort.Logout(dtos.LogoutDTO{
		Token: request.Token,
	})
	utils.Success(ctx, http.StatusNoContent, "Success", nil)
}
