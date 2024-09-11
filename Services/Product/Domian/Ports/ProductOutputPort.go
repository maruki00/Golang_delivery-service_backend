package product_domain_ports

import shared_models "delivery/Services/Shared/Infrastructure/Models"

type ProductOutputPort interface {
}

func NewProductOutputPort() JsonProductOutputPort {
	return JsonProductOutputPort{}
}

func (obj *JsonProductOutputPort) Success(data shared_models.ResponseModel) shared_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}

func (obj *JSONAuthPresenter) Error(data shared_models.ResponseModel) domain_auth_contracts.ViewModel {
	return shared_adapters_viewmodels.NewJsonViewModel(data)
}
