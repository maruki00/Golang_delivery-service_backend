package shared_core

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (o *RequestValidator) Validate(ctx gin.Context, request *interface{}) {

}
