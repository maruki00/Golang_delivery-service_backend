package product_domain_ports

import (
	shared_models "delivery/Services/Shared/Infrastructure/Models"
	shared_adapters_viewmodels "delivery/Services/Shared/Usergateway/Adapters/ViewModels"
)

type ProductOutputPort interface {
}

func NewProductOutputPort() shared_ {
	return JsonProductOutputPort{}
}

func (obj *JsonProductOutputPort) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *JsonProductOutputPort) Error(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
