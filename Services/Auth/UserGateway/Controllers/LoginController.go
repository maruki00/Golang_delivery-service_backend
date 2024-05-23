package Authcontrollers

import (
	Authusecases "delivery/Services/Auth/Application/UseCases"
	"delivery/Services/Auth/Domain/DTOs"
	SharedUtils "delivery/Services/Sahared/Application/Utils"

	"github.com/gin-gonic/gin"
)

type Authcontroller struct {
}

var (
	userCase = &Authusecases.LoginuserCase{}
)

func Login(ctx *gin.Context) {
	var dto DTOs.LoginDTO
	SharedUtils.ParseBody(ctx.Request, &dto)
	res := userCase.Login(dto)
	SharedUtils.Success(ctx, res, 200)
}

func Register(ctx *gin.Context) {

	var dto DTOs.UserDTO
	SharedUtils.ParseBody(ctx.Request, &dto)
	res := userCase.Register(dto)
	SharedUtils.Success(ctx, res, 200)
}
