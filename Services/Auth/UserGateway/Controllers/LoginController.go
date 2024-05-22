package Authcontrollers

import (
	Authusecases "delivery/Services/Auth/Application/UseCases"
	"delivery/Services/Auth/Domain/DTOs"
	SharedUtils "delivery/Services/Sahared/Application/Utils"

	"github.com/gin-gonic/gin"
)

type Authcontroller struct {
}

func (obj *Authcontroller) Login(ctx *gin.Context) {
	var dto *DTOs.LoginDTO
	SharedUtils.ParseBody(ctx.Request, dto)
	Authusecases.LoginuserCase(dto)
	SharedUtils.Success(ctx, dto, 200)
}
