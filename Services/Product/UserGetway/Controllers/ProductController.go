package product_usergetway_controllers

import (
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

	request := &product_usergetway_requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.inPort.Insert(&product_domain_dtos.InsertProductDTO{
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "someting wrong happen")
		return
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{"product": res})
}

func (obj *ProductController) Search(ctx *gin.Context) {

	request := &product_usergetway_requests.SearchProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.inPort.Search(&product_domain_dtos.SearchProductDTO{
		Query: request.Query,
	})

	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "someting wrong happen")
		return
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{"products": res})
}

func (obj *ProductController) Update(ctx *gin.Context) {

	request := &product_usergetway_requests.UpdateProductRequerst{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "validation : "+err.Error())
		return
	}
	res, err := obj.inPort.Update(&product_domain_dtos.UpdateProductDTO{
		Id:    request.Id,
		Label: request.Label,
		Type:  request.Type,
		Price: request.Price,
	})
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "someting wrong happen "+err.Error())
		return
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{"product": res})
}

func (obj *ProductController) Delete(ctx *gin.Context) {

	request := &product_usergetway_requests.DeleteProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.inPort.Delete(&product_domain_dtos.DeleteProductDTO{
		Id: request.Id,
	})
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "someting wrong happen")
		return
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{"product": res})
}

func (obj *ProductController) GetProduct(ctx *gin.Context) {
	request := &product_usergetway_requests.GetProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.inPort.GetById(&product_domain_dtos.GetProductByIdDTO{
		Id: request.Id,
	})

	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", "someting wrong happen")
		return
	}

	if res.GetId() == 0 {
		res = nil
	}

	shared_utils.Success(ctx, http.StatusOK, "Success", gin.H{"product": res})
}
