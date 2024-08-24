package product_usergetway_controllers

import (
	product_services "delivery/Services/Product/Application/Services"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	product_Infrastructure_repository "delivery/Services/Product/Infrastructure/Repositories"
	product_usergetway_requests "delivery/Services/Product/UserGetway/Requests"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductController struct {
	service   *product_services.ProductService
	Validator *validator.Validate
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		service: product_services.NewProductService(product_Infrastructure_repository.NewProductRepository(db, &product_infrastructure_models.Product{})),
	}
}

func (obj *ProductController) Insert(ctx *gin.Context) {

	var request product_usergetway_requests.InsertProductRequest

	obj.Validator.Struct()
	fmt.Println("Insert New Product ...")
}
