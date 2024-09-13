package product_services

import (
	product_domain_dtos "delivery/Services/Product/Domian/DTOS"
	product_domain_ports "delivery/Services/Product/Domian/Ports"
	product_domain_repositories "delivery/Services/Product/Domian/Repositories"
	product_infrastructure_models "delivery/Services/Product/Infrastructure/Models"
	shared_domain_contracts "delivery/Services/Shared/Domain/Contracts"
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	"net/http"

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

func (obj *ProductService) Insert(dto *product_domain_dtos.InsertProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Insert(&product_infrastructure_models.Product{
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

func (obj *ProductService) Search(dto *product_domain_dtos.SearchProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Search(dto.Query)
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

func (obj *ProductService) Update(dto *product_domain_dtos.UpdateProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Update(dto.Id, map[string]interface{}{
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

func (obj *ProductService) Delete(dto *product_domain_dtos.DeleteProductDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.Delete(dto.Id)

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

func (obj *ProductService) GetById(dto *product_domain_dtos.GetProductByIdDTO) shared_domain_contracts.ViewModel {

	res, err := obj.productRepository.GetById(dto.Id)
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
