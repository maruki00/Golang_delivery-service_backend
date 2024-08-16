package auth_usergetway_controllers

import (
	auth_requests "delivery/Services/Auth/UserGateway/Requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type AuthController struct {
	Validate *validator.Validate
}

func NewAuthController() *AuthController {
	return &AuthController{
		Validate: validator.New(),
	}
}

func (obj AuthController) Login(ctx *gin.Context) {

	request := &auth_requests.LoginRequest{}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		fmt.Println(err.Error())
		// 	validationErrors := err.(validator.ValidationErrors)
		// 	errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		// 	return
		// }

		// res, _ := authu_services.Login(auth_domain_dtos.LoginDTO{
		// 	Login:    request.Login,
		// 	Password: request.Password,
	}
	//)

	ctx.JSON(200, request)
}

func Register(ctx *gin.Context) {
	// var dto DTOs.UserDTO
	// SharedUtils.ParseBody(ctx.Request, &dto)
	// fmt.Println("level Register Controller : ", ctx.Request.Body)
	// res := userCase.Register(dto)
	// SharedUtils.Success(ctx, res, 200)
}
