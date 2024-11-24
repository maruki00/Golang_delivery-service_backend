package order_usergateway_controllers

import (
	order_domain_ports "delivery/Internal/Order/Domain/Ports"
	order_usergateway_requests "delivery/Internal/Order/UserGateway/Requests"
	shared_models "delivery/Internal/Shared/Infrastructure/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/net/context"
)

type OrderController struct {
	context   context.Context
	inputPort order_domain_ports.OrderInputPort
	Validate  *validator.Validate
}

func NewOrderController(inputPort order_domain_ports.OrderInputPort) *OrderController {
	ctx, cancleFunc := context.WithCancel(context.Background())
	defer cancleFunc()
	return &OrderController{
		context:   ctx,
		inputPort: inputPort,
		Validate:  validator.New(),
	}
}

func (obj *OrderController) Create(ctx *gin.Context) {
	request := &order_usergateway_requests.CreateOrderRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(http.StatusBadRequest, &shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Error:   "invalid request params",
			Result:  nil,
			Message: "request error",
		})
		return
	}

	if err := obj.Validate.Struct(request); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		ctx.JSON(http.StatusBadRequest, &shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Error:   errorMessage,
			Result:  nil,
			Message: "request error",
		})
		return
	}
}
