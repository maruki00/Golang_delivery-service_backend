package order_usergateway_controllers

import (
	order_domain_ports "delivery/Services/Order/Domain/Ports"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/net/context"
)

type OrderController struct {
	context   context.Context
	inputPort order_domain_ports.OrderInputPort
	Validate  *validator.Validate
}

func NewOrderController(context context.Context, inputPort order_domain_ports.OrderInputPort) *OrderController {
	return &OrderController{
		context:   context,
		inputPort: inputPort,
		Validate:  validator.New(),
	}
}


func (obj*OrderController) Create(ctx *gin.Context) {
	var 
}