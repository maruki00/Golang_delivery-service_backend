package product_usergetway_controllers

import (
	"context"
	product_services "delivery/internal/product/Application/Services"
	product_domain_dtos "delivery/internal/product/Domian/DTOS"
	product_domain_ports "delivery/internal/product/Domian/Ports"
	product_usergateway_adapters_presenters "delivery/internal/product/UserGetway/Adapters/Presenters"
	product_usergateway_requests "delivery/internal/product/UserGetway/Requests"
	product_infra_models "delivery/internal/product/infra/Models"
	product_infra_repository "delivery/internal/product/infra/Repositories"
	shared_utils "delivery/internal/shared/Application/Utils"
	shared_core "delivery/internal/shared/infra/Core"
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
			product_infra_repository.NewProductRepository(db, &product_infra_models.Product{}),
			&product_usergateway_adapters_presenters.ProductPresenter{},
		),
		Validate: validator.New(),
	}
}

func (obj *ProductController) Insert(ctx *gin.Context) {

	c, Cancel := context.WithCancel(context.Background())
	defer Cancel()
	request := &product_usergateway_requests.InsertProductRequest{}

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
	request := &product_usergateway_requests.SearchProductRequest{}

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
	request := &product_usergateway_requests.UpdateProductRequerst{}

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

	request := &product_usergateway_requests.DeleteProductRequest{}

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

	request := &product_usergateway_requests.GetProductRequest{}

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
	request := &product_usergateway_requests.MultipleProductstRequest{}

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
