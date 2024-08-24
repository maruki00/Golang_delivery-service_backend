package product_usergetway_controllers

import (
	product_services "delivery/Services/Product/Application/Services"
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	product_Infrastructure_repository "delivery/Services/Product/Infrastructure/Repositories"
	product_usergetway_requests "delivery/Services/Product/UserGetway/Requests"
	shared_utils "delivery/Services/Shared/Application/Utils"
	shared_core "delivery/Services/Shared/Infrastructure/Core"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductController struct {
	service  *product_services.ProductService
	Validate *validator.Validate
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		service:  product_services.NewProductService(product_Infrastructure_repository.NewProductRepository(db, &product_infrastructure_models.Product{})),
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
	res, err := obj.service.Insert(&product_domain_dtos.InsertProductDTO{
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

	request := &product_usergetway_requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.service.Insert(&product_domain_dtos.InsertProductDTO{
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


func (obj *ProductController) Update(ctx *gin.Context) {

	request := &product_usergetway_requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.service.Insert(&product_domain_dtos.InsertProductDTO{
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


func (obj *ProductController) Delete(ctx *gin.Context) {

	request := &product_usergetway_requests.InsertProductRequest{}

	err := shared_core.Validate(ctx, obj.Validate, request)
	if err != nil {
		shared_utils.Error(ctx, http.StatusBadRequest, "Error", err.Error())
		return
	}
	res, err := obj.service.Insert(&product_domain_dtos.InsertProductDTO{
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


