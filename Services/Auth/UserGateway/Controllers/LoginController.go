package auth_usergetway_controllers

import (
	authu_services "delivery/Services/Auth/Application/Services"
	auth_domain_dtos "delivery/Services/Auth/Domain/DTOs"
	SharedUtils "delivery/Services/Shared/Application/Utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService authu_services.AuthService
}

func (obj AuthController) Login(ctx *gin.Context) {
	var dto auth_domain_dtos.LoginDTO
	SharedUtils.ParseBody(ctx.Request, &dto)
	fmt.Println("body : ", dto, ctx.Request)
	res, _ := obj.authService.Login(dto)
	SharedUtils.Success(ctx, res, 200)
}

func Register(ctx *gin.Context) {
	// var dto DTOs.UserDTO
	// SharedUtils.ParseBody(ctx.Request, &dto)
	// fmt.Println("level Register Controller : ", ctx.Request.Body)
	// res := userCase.Register(dto)
	// SharedUtils.Success(ctx, res, 200)
}
