package product_usergetway_controllers

import (
	"context"
	product_services "delivery/Services/Product/Application/Services"
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_ports "delivery/Services/Product/Domian/Ports"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	product_Infrastructure_repository "delivery/Services/Product/Infrastructure/Repositories"
	product_usergateway_adapters_presenters "delivery/Services/Product/UserGetway/Adapters/Presenters"
	product_usergetway_requests "delivery/Services/Product/UserGetway/Requests"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_core "delivery/Services/Shared/Infrastructure/Core"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductController struct {
	Validate *validator.Validate
	inPort   product_domain_ports.ProductInputPort
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		inPort: product_services.NewProductService(
			product_Infrastructure_repository.NewProductRepository(db, &product_infrastructure_models.Product{}),
			&product_usergateway_adapters_presenters.ProductPresenter{},
		),
		Validate: validator.New(),
	}
}

func (obj *ProductController) Insert(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &product_usergetway_requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Insert(c, &product_domain_dtos.InsertProductDTO{
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Search(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &product_usergetway_requests.SearchProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Search(c, &product_domain_dtos.SearchProductDTO{
		Query: request.Query,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Update(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &product_usergetway_requests.UpdateProductRequerst{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "validation : "+err.Error())
		return
	}
	res := obj.inPort.Update(c, &product_domain_dtos.UpdateProductDTO{
		Id:    request.Id,
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) Delete(ctx *gin.Context) {

	// c, Cancel := context.WithCancel(context.Background())
	c, Cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	defer Cancel()

	request := &product_usergetway_requests.DeleteProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.Delete(c, &product_domain_dtos.DeleteProductDTO{
		Id: request.Id,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) GetProduct(ctx *gin.Context) {
	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()

	request := &product_usergetway_requests.GetProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.GetById(c, &product_domain_dtos.GetProductByIdDTO{
		Id: request.Id,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}

func (obj *ProductController) MultipleProducts(ctx *gin.Context) {
	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &product_usergetway_requests.MultipleProductstRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res := obj.inPort.MultipleProducts(c, &product_domain_dtos.MultipleProductsDTO{
		Ids: request.Ids,
	})

	ctx.JSON(res.GetResponse().Status, res.GetResponse())
}
