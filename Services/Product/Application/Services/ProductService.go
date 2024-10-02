package product_services

import (
	"context"
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_ports "delivery/Services/Product/Domian/Ports"
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	productRepository product_domain_repositories.IProductRepository
	outputPort        product_domain_ports.ProductOutputPort
}

func NewProductService(productRepository product_domain_repositories.IProductRepository, outputPort product_domain_ports.ProductOutputPort) *ProductService {
	return &ProductService{
		productRepository: productRepository,
		outputPort:        outputPort,
	}
}

func (obj *ProductService) Insert(ctx context.Context, dto *product_domain_dtos.InsertProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Insert(ctx, &product_infrastructure_models.Product{
		Label: dto.Label,
		Type:  dto.Type,
		Price: dto.Price,
	})
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) Search(ctx context.Context, dto *product_domain_dtos.SearchProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Search(ctx, dto.Query)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"products": res},
	})
}

func (obj *ProductService) Update(ctx context.Context, dto *product_domain_dtos.UpdateProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Update(ctx, dto.Id, map[string]interface{}{
		"id":    dto.Id,
		"label": dto.Label,
		"type":  dto.Type,
		"price": dto.Price,
	})

	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) Delete(ctx context.Context, dto *product_domain_dtos.DeleteProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Delete(ctx, dto.Id)

	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not delete the product : " + err.Error(),
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) GetById(ctx context.Context, dto *product_domain_dtos.GetProductByIdDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.GetById(ctx, dto.Id)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not get the product or doesnt exsits",
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"product": res},
	})
}

func (obj *ProductService) MultipleProducts(ctx context.Context, dto *product_domain_dtos.MultipleProductsDTO) shared_domain_contracts.ViewModel {

	ids := []int{}
	strIds := strings.Split(dto.Ids, ",")

	for _, id := range strIds {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		ids = append(ids, int(i))
	}

	res, err := obj.productRepository.GetProductByMultipleId(ctx, ids)
	if err != nil {
		return obj.outputPort.Error(shared_models.ResponseModel{
			Status:  http.StatusBadRequest,
			Message: "Error",
			Error:   "could not get the product or doesnt exsits",
			Result:  nil,
		})
	}

	return obj.outputPort.Success(shared_models.ResponseModel{
		Status:  http.StatusOK,
		Message: "Success",
		Error:   nil,
		Result:  gin.H{"products": res},
	})
}
