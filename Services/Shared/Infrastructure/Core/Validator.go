package shared_core

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (o *RequestValidator) Validate(ctx gin.Context, request interface{}) (error, interface{}) {

	if err := ctx.BindJSON(&request); err != nil {
		return err, nil
	}

	if err := o.Validator.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return fmt.Errorf("Validation failed for field: %s", validationErrors[0].Field()), nil
	}
	return nil, request
}
